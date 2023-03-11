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

type Test struct {
	Testval string `json:"testval"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Test `json:"data"`
	Message string `json:"message"`
}

func GetTest(c *gin.Context) {
	db := setupDB()
	rows, err := db.Query("SELECT * FROM pgtest")
	checkErr(err)

	var tests []Test

	for rows.Next() {
		var testval string
		err = rows.Scan(&testval)
		checkErr(err)
		tests = append(tests, Test{Testval: testval})
	}

	var response = JsonResponse{Type: "success", Data: tests}

	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", GetTest)

	r.Run()
}
