package main

import (
	"log"
	"time"

	"github.com/KirkPig/paintplz-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	var database = NewSQLConn()

	defer database.Close()

	router := gin.Default()
	v1paintPlz := router.Group("api/paintplz/v1")
	{
		v1paintPlz.POST("/register", services.RegisterHandler)
		v1paintPlz.POST("/login", services.LoginHandler)
		v1paintPlz.POST("/search_artist", services.SearchArtistHandler)
		v1paintPlz.GET("/artist_profile/:user_id", services.GetArtistProfileHandler)
		v1paintPlz.POST("/artist_profile/artwork/upload", services.UploadArtworkHandler)
		v1paintPlz.POST("/artist_profile/artwork/edit", services.EditArtworkHandler)
		v1paintPlz.POST("/artist_profile/artwork/delete", services.DeleteArtworkHandler)
		v1paintPlz.GET("/tags", services.GetTagsHandler)
	}

	router.Run("localhost:1323")

}

func NewSQLConn() *gorm.DB {

	conf := mysql.Config{
		DBName: "PAINTPLZIO",
		User:   "root",
		Passwd: "123456",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		Loc:    time.Local,
	}

	conn, err := gorm.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatalln("connection error")
	}

	return conn

}
