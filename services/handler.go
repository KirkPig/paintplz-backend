package services

import (
	"log"
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
		return
	}
	log.Println(req)
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
		return
	}
	response, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) SearchArtistHandler(c *gin.Context) {

	var req SearchArtistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
		return
	}
	if req.MaxPriceRate == nil {
		*req.MaxPriceRate = 1000000000000
	}
	if req.MinPriceRate == nil {
		*req.MinPriceRate = 0
	}
	if req.MaxRating == nil {
		*req.MaxRating = 10
	}
	if req.MinRating == nil {
		*req.MinRating = 0
	}
	result, err := h.service.SearchAritst(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetArtistProfileHandler(c *gin.Context) {

}

func (h *Handler) UploadArtworkHandler(c *gin.Context) {

	var req UploadArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
		return
	}
	err := h.service.UploadArtwork(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successfully",
	})
}

func (h *Handler) EditArtworkHandler(c *gin.Context) {

	var req EditArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
		return
	}

	err := h.service.EditArtwork(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successfully",
	})

}

func (h *Handler) DeleteArtworkHandler(c *gin.Context) {

	var req DeleteArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
		return
	}
	err := h.service.DeleteArtwork(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successfully",
	})

}

func (h *Handler) GetTagsHandler(c *gin.Context) {

}
