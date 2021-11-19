package main

import (
	"database/sql"

	"github.com/KirkPig/paintplz-backend/services"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	sqlDB, err := sql.Open("mysql", "paintplzuser:password@/paintplz")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err.Error())
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	new_uuid, err := uuid.NewV4()
	if err != nil {
		panic(err.Error())
	}

	gormDB.Raw("call Register(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", new_uuid.String(), "Man", "Name", "Surname", "Man@gmail.com",
		"12321351513", "12345678", "artist", 100, 230, "my name is ...")
	/**
	IN PAINTPLZUSERID VARCHAR(100),
	IN USERNAME VARCHAR(30),
	IN NAME VARCHAR(30),
	IN SURNAME VARCHAR(30),
	IN EMAIL VARCHAR(30),
	IN CITIZENID VARCHAR(13),
	IN PASSWORD VARCHAR(64),
	IN USERTYPE ENUM('artist','customer'),
	IN MINIMUMPRICERATE FLOAT(12,2),
	IN MAXIMUMPRICERATE FLOAT(12,2),
	IN BIOGRAPHY TEXT(500)

	*/
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
