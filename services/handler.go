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
	if req.MinPriceRate == nil {
		req.MinPriceRate = new(int64)
		*req.MinPriceRate = 0
	}
	if req.MaxPriceRate == nil {
		req.MaxPriceRate = new(int64)
		*req.MaxPriceRate = 1000000
	}
	if req.MinRating == nil {
		req.MinRating = new(float32)
		*req.MinRating = 0.0
	}
	if req.MaxRating == nil {
		req.MaxRating = new(float32)
		*req.MaxRating = 5.1
	}

	result, err := h.service.SearchArtist(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetArtistProfileHandler(c *gin.Context) {

	user_id := c.Param("user_id")

	log.Println(user_id)

	response, err := h.service.ArtistProfile(user_id)

	log.Println(response)

	if err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Response failed",
		})
		return
	}

	c.JSON(http.StatusOK, response)

}

func (h *Handler) UploadArtworkHandler(c *gin.Context) {

	var req UploadArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
		return
	}
	result, err := h.service.UploadArtwork(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) EditArtworkHandler(c *gin.Context) {

	var req EditArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
		return
	}

	log.Println(req)

	result, err := h.service.EditArtwork(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)

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
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete successfully",
	})

}

func (h *Handler) GetTagsHandler(c *gin.Context) {

	tags, err := h.service.GetAllTag()

	if err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Failed tp fetch",
		})
		return
	}

	c.JSON(http.StatusOK, GetTagResponse{
		Tags: tags,
	})

}

func (h *Handler) UploadArtworkMongo(c *gin.Context) {
	var req UploadArtworkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Logger{
			Log: "Handler error",
		})
	}
	artwork, err := h.service.UploadArtworkMongo(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		},
		)
	}
	c.JSON(http.StatusOK, artwork)
}
