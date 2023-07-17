package controller

import (
	"character_search/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var characters = []model.Character{
	{Name: "アドマイヤベガ"},
	{Name: "キングヘイロー"},
	{Name: "ナリタブライアン"},
}

func GetCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters)
}
