package main

import (
	"github.com/KirkPig/paintplz-backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
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
