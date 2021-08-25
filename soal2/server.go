package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"soal2/config"
	"soal2/route"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

var err error

func main() {
	godotenv.Load()
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}
	e := echo.New()
	route.GenRoutes(e)
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":1323"))
}
