package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var characters = []Character{
	{ID: 1, Name: "test"},
	{ID: 2, Name: "test2"},
}

func main() {
	router := gin.Default()

	router.GET("/characters", getCharacters)

	router.Run()
}

func getCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters)
}
