package readDB

import (
	"encoding/json"
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

type RoleWolf struct {
	Role int  `json:"role"`
	Wolf bool `json:"wolf"`
}

func GetRoleWolf(playerId string) RoleWolf {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Println("error getting Player document: \n", err)

	}
	iter, err := client.Collection("Player").Doc(playerId).Get(ctx)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}

	bytes, _ := json.Marshal(iter.Data()["Role"])
	var roleInt int64
	err = json.Unmarshal(bytes, &roleInt)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}

	var roleWolf RoleWolf
	roleWolf.Role = int(roleInt)
	roleWolf.Wolf = iter.Data()["Wolf"].(bool)
	defer client.Close()

	return roleWolf
}
