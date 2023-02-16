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

func main() {
	// インスタンスを作成
	e := echo.New()
	e.Use(middleware.CORS())

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	// ローカル環境の場合、http://localhost:1323/
	e.GET("/is-room-exit/:id/:password", isRoomExit)
	e.GET("/partic-list/:roomId", func(c echo.Context) error {
		playerList := getParticList(c)
		return c.JSON(http.StatusOK, playerList)
	})
	e.GET("/theme:description", getTheme)
	e.GET("/hint-list/:roomId", func(c echo.Context) error {
		hintList := getHintList(c)
		return c.JSON(http.StatusOK, hintList)
	})
	e.GET("/step/:roomId", getStep)
	e.GET("/random-theme", getRandomTheme)
	e.POST("/createRoom", createRoom)
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

func isRoomExit(c echo.Context) error {
	var exit bool = true
	id := c.Param("id")
	password := c.Param("password")

	fmt.Println(id, password) //test
	return c.JSON(http.StatusOK, exit)
}

func getParticList(c echo.Context) [][]interface{} {
	var playerList = [][]interface{}{
		{"tanaka", 1},
		{"suzuki", 2},
		{"mashio", 3},
	}
	id := c.Param("roomId")
	fmt.Println(id) //test
	return playerList
}

func getTheme(c echo.Context) error {
	var theme string = "テスト"
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
	var step int = 1
	return c.JSON(http.StatusOK, step)
}
func getRandomTheme(c echo.Context) error {
	var theme string = "テスト"
	return c.JSON(http.StatusOK, theme)
}

func createRoom(c echo.Context) error {
	reqBody := new(Room)s
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	password := reqBody.Password
	particNum := reqBody.PaticNum

	useDB.CreateRoom(password, particNum, "theme", 0, 0)

	return c.String(http.StatusOK, "OK")
}

//$body = @{
//     password = "mypassword"
//     particNum = 5
// } | ConvertTo-Json

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/createRoom -Body $body -ContentType "application/json"