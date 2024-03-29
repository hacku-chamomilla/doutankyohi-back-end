package createDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	readDB "gamename-back-end/pkg/cruds/read"

	"cloud.google.com/go/firestore"
)

func DecideTheme(roomId string, howToDecideTheme int) bool {

	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Set(ctx, map[string]interface{}{
		"HowToDecideTheme": howToDecideTheme,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}

	var step int = 2
	if howToDecideTheme == 1 {
		roomRef := client.Collection("Room").Doc(roomId)
		_, err = roomRef.Set(ctx, map[string]interface{}{
			"Theme": readDB.GetRandomTheme(),
		}, firestore.MergeAll)
		if err != nil {
			return false
		}
		step = 3
	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: step},
	})
	if err != nil {
		return false
	}
	defer client.Close()
	return true
}
