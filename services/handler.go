package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func RegisterHandler(c *gin.Context) {

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func LoginHandler(c *gin.Context) {

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func SearchArtistHandler(c *gin.Context) {

	var req SearchArtistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func GetArtistProfileHandler(c *gin.Context) {

}

func UploadArtworkHandler(c *gin.Context) {

	var req UploadArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func EditArtworkHandler(c *gin.Context) {

	var req EditArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func DeleteArtworkHandler(c *gin.Context) {

	var req DeleteArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func GetTagsHandler(c *gin.Context) {

}
