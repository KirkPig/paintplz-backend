package main

type RegisterRequest struct {
	Username     string `json:"username" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
	Email        string `json:"email" binding:"required"`
	CitizenID    string `json:"citizenID" binding:"required"`
	Password     string `json:"password" binding:"required"`
	UserType     bool   `json:"userType" binding:"required"`
	MinPriceRate int64  `json:"minPriceRate"`
	MaxPriceRate int64  `json:"maxPriceRate"`
	Biography    string `json:"biography"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SearchArtistRequst struct {
	ArtistName   string  `json:"artistName" binding:"required"`
	MinPriceRate int64   `json:"minPriceRate"`
	MaxPriceRate int64   `json:"maxPriceRate"`
	MinRating    float32 `json:"minRating"`
	MaxRating    float32 `json:"maxRating"`
	Tags         []Tag   `json:"tags"`
}

type Tag struct {
	TagId   string `json:"tagID" binding:"required"`
	TagName string `json:"tagName" binding:"required"`
}

type SearchResultResponse struct {
	Username string  `json:"username" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Surname  string  `json:"surname" binding:"required"`
	Rating   float32 `json:"rating"`
}
