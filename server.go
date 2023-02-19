package main

import (
	"fmt"
	useDB "gamename-back-end/src/use_DB"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Room struct { //TODO　create_dbと被るからそこを考えよう
	Password string `json:"password"`
	PaticNum int    `json:"particNum"`
}

type Player struct {
	RoomId     string `json:"roomId"`
	PlayerName string `json:"playerName"`
	PlayerIcon int    `json:"playerIcon"`
}

type GetTheme struct {
	PlayerId string `json:"playerId"`
	Theme    string `json:"theme"`
}

type GetHint struct {
	PlayerId string `json:"playerId"`
	Hint     string `json:"hint"`
}

type DeleteHint struct {
	Hint []string `json:"hint"`
}

func main() {
	// インスタンスを作成
	e := echo.New()
	e.Use(middleware.CORS())

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	// ローカル環境の場合、http://localhost:1323/
	e.GET("/is-room-exit/", isRoomExit)
	e.GET("/partic-list/", func(c echo.Context) error {
		playerList := getParticList(c)
		return c.JSON(http.StatusOK, playerList)
	})
	e.GET("/theme", getTheme)
	e.GET("/hint-list/:roomId", func(c echo.Context) error {
		hintList := getHintList(c)
		return c.JSON(http.StatusOK, hintList)
	})
	e.GET("/step", getStep)
	e.GET("/random-theme", getRandomTheme)
	e.POST("/createRoom", createRoom)
	e.POST("/addPlayer", postAddPlayer)
	e.POST("/createTheme", postCreateTheme)
	e.POST("/createHint", postCreateHint)
	e.POST("/deleteHint", postDeleteHint)
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

func isRoomExit(c echo.Context) error {
	var exit bool = true
	rid := c.QueryParam("rid")
	password := c.QueryParam("password")

	useDB.IsRoomExit(rid, password)
	return c.JSON(http.StatusOK, exit)
}

func getParticList(c echo.Context) [][]interface{} {
	roomid := c.QueryParam("rid")
	fmt.Println(roomid)
	playerList := useDB.PlayerList(roomid) //test]
	fmt.Println(playerList)
	return playerList
}

func getTheme(c echo.Context) error {
	roomid := c.QueryParam("roomid")
	theme := useDB.GetTheme(roomid)
	return c.JSON(http.StatusOK, theme)
}

func getHintList(c echo.Context) [][]interface{} {
	var hintList = [][]interface{}{
		{"key", "hint1", true},
		{"key2", "hint2", true},
		{"key3", "hint3", true},
	}
	return hintList
}
func getStep(c echo.Context) error {
	roomid := c.QueryParam("roomid")
	fmt.Println(roomid)
	return c.JSON(http.StatusOK, useDB.GetStep(roomid))
}
func getRandomTheme(c echo.Context) error {
	var theme string = "テスト"
	return c.JSON(http.StatusOK, theme)
}

func createRoom(c echo.Context) error {
	reqBody := new(Room)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	password := reqBody.Password
	particNum := reqBody.PaticNum

	useDB.CreateRoom(password, particNum, "theme", 0, 0)

	return c.String(http.StatusOK, "OK")
}

func postAddPlayer(c echo.Context) error {
	reqBody := new(Player)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	roomId := reqBody.RoomId
	playerName := reqBody.PlayerName
	playerIcon := reqBody.PlayerIcon

	fmt.Println(roomId, playerName, playerIcon)
	playerId := useDB.AddPlayer(roomId, playerName, playerIcon)
	fmt.Println(playerId)
	return c.JSON(http.StatusOK, playerId)
}

func postCreateTheme(c echo.Context) error {
	reqBody := new(GetTheme)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	theme := reqBody.Theme

	return c.JSON(http.StatusOK, useDB.CreateTheme(theme, playerId))
}
func postCreateHint(c echo.Context) error {
	reqBody := new(GetHint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	playerId := reqBody.PlayerId
	hint := reqBody.Hint
	fmt.Println(hint)
	return c.JSON(http.StatusOK, useDB.CreateHint(hint, playerId))
}
func postDeleteHint(c echo.Context) error {
	reqBody := new(DeleteHint)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	hintList := reqBody.Hint
	return c.JSON(http.StatusOK, useDB.DeleteHint(hintList))
}

// $body = @{
//     password = "yourpass"
//     particNum = 3
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/createRoom -Body $body -ContentType "application/json"
//curl -d "roomId = cbBipgOwuA8wxu5XAXFW" -d "playerName = testman" -d "playerIcon = 3" http://localhost:1323/addPlayer
