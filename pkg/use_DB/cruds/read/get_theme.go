package useDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func GetTheme(roomId string) string {
	ctx, client, _err := connectDB.ConnectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Room documents: %v\n", err)
	}
	theme := iter.Data()["Theme"].(string)
	defer client.Close()
	return theme

}