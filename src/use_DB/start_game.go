package useDB

import (
	"math/rand"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
)

func StartGame(roomId string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Update(ctx, []firestore.Update{
		{Path: "Phase", Value: firestore.Increment(1)},
	})

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()

	if err != nil {
		return false
	}

	var playerList []string

	for _, rpDoc := range rpDocs {
		playerId := rpDoc.Data()["PlayerId"].(string)
		playerList = append(playerList, playerId)
	}

	for i := range playerList {
		j := rand.Intn(i + 1)
		playerList[i], playerList[j] = playerList[j], playerList[i]
	}

	var roleCount int = 1

	for i := range playerList {
		playerDoc := client.Collection("Player").Doc(playerList[i])
		if err != nil {
			return false
		}
		_, _err := playerDoc.Set(ctx, map[string]interface{}{
			"Role": strconv.Itoa(roleCount),
		}, firestore.MergeAll)

		if _err != nil {
			return false
		}
		if roleCount != 3 {
			roleCount += 1
		}
	}

	rand.Seed(time.Now().UnixNano())
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"Step": 1,
	}, firestore.MergeAll)
	if _err != nil {
		return false
	}
	defer client.Close()
	return true
}
