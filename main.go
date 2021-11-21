package main

import (
	"log"
	"time"

	"github.com/KirkPig/paintplz-backend/repository"
	"github.com/KirkPig/paintplz-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	var database = NewSQLConn()

	defer database.Close()

	//mongo_repository.Init()
	//// sample use of mongodb
	/*
		ctx, cancel := repository.GetContext()
		artwork := repository.ArtworkMongo{
			ArtworkID:   "65432132131",
			Title:       "my first artwork",
			Description: "strongest artwork",
			UploadDate:  "1-1-2021",
			ArtworkUrl:  "https://yt3.ggpht.com/ytc/AKedOLRnmGhH8h_Le-nLsSZwx7K0pgRA0Zdr8bU9XgXkkw=s900-c-k-c0x00ffffff-no-rj",
			Tags: []repository.TagsMongo{
				{
					TagID:   "1",
					TagName: "3d",
				},
				{
					TagID:   "5",
					TagName: "4d",
				},
			},
		}
		_, err := repository.ArtworkCollection.InsertOne(ctx, artwork)
		defer cancel()
		if err != nil {
			log.Println("What happend " + err.Error())
		}*/
	//	repository.ArtworkCollection.InsertOne(c)

	handler := services.NewHandler(*services.NewService(*repository.New(database)))
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	v1paintPlz := router.Group("api/paintplz/v1")
	{
		v1paintPlz.POST("/register", handler.RegisterHandler)
		v1paintPlz.POST("/login", handler.LoginHandler)
		v1paintPlz.POST("/search_artist", handler.SearchArtistHandler)
		v1paintPlz.GET("/artist_profile/:user_id", handler.GetArtistProfileHandler)
		v1paintPlz.POST("/artist_profile/artwork/upload", handler.UploadArtworkHandler)
		v1paintPlz.POST("/artist_profile/artwork/edit", handler.EditArtworkHandler)
		v1paintPlz.POST("/artist_profile/artwork/delete", handler.DeleteArtworkHandler)
		v1paintPlz.GET("/tags", handler.GetTagsHandler)
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
