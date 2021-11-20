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

func (h *Handler) RegisterHandler(c *gin.Context) {

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}
	err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "registered successfully",
	})
}

func (h *Handler) LoginHandler(c *gin.Context) {

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func (h *Handler) SearchArtistHandler(c *gin.Context) {

	var req SearchArtistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func (h *Handler) GetArtistProfileHandler(c *gin.Context) {

}

func (h *Handler) UploadArtworkHandler(c *gin.Context) {

	var req UploadArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func (h *Handler) EditArtworkHandler(c *gin.Context) {

	var req EditArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func (h *Handler) DeleteArtworkHandler(c *gin.Context) {

	var req DeleteArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}

}

func (h *Handler) GetTagsHandler(c *gin.Context) {

}
