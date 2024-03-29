package createDB

import (
	"encoding/json"
	connectDB "gamename-back-end/pkg/connect_db"
	"log"

	"cloud.google.com/go/firestore"
)

func CreateHint(inputHint string, playerId string, roomId string) bool {

	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	docRef := client.Collection("Player").Doc(playerId)
	_, nil := docRef.Set(ctx, map[string]interface{}{
		"Hint": inputHint,
	}, firestore.MergeAll)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()

	var addStep bool = true

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			return false
		}
		bytes, _ := json.Marshal(playerDoc.Data()["Role"])
		var roleInt int64
		err = json.Unmarshal(bytes, &roleInt)
		if err != nil {
			return false
		}
		if int(roleInt) != 1 {
			if playerDoc.Data()["Hint"].(string) == "no-hint" {
				addStep = false
			}
		}

	}

	if addStep == true {
		_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
			{Path: "Step", Value: 4},
		})
		if err != nil {
			return false
		}
	}

	defer client.Close()
	return true
}
