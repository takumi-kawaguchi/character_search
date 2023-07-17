package main

import (
	"character_search/controller"
	"character_search/model"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load environment file")
	}

	host := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, password, host, dbName)
	fmt.Println(host)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Character{})
	db.Create(&model.Character{Name: "アドマイヤベガ"})

	router := gin.Default()
	router.GET("/characters", controller.GetCharacters)
	router.Run()
}
