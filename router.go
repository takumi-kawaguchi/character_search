package main

import (
	"character_search/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.GET("/characters", controller.GetCharacters)
	r.Run()
}
