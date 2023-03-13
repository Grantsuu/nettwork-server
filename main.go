package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Ball struct {
	Name string
}

const (
	DB_USER     = "grantsherman"
	DB_PASSWORD = "akkomari"
	DB_NAME     = "nettwork"
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Game struct {
	Id          int    `json:"testval"`
	Owner       string `json:"owner"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	Variant     string `json:"variant"`
	Max         int    `json:"max"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Data    []Game `json:"data"`
	Message string `json:"message"`
}

func GetGames(c *gin.Context) {
	db := setupDB()
	rows, err := db.Query("SELECT * FROM games")
	checkErr(err)

	var games []Game

	for rows.Next() {
		var id int
		var owner string
		var title string
		var description string
		var date string
		var location string
		var variant string
		var max int
		err = rows.Scan(&id, &owner, &title, &description, &date, &location, &variant, &max)
		checkErr(err)
		games = append(
			games,
			Game{Id: id, Owner: owner, Title: title, Description: description, Date: date, Location: location, Variant: variant, Max: max})
	}

	var response = JsonResponse{Status: "success", Data: games}

	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/games", GetGames)

	r.Run()
}
