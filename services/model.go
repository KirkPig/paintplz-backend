package services

type UploadArtworkRequest struct {
	Username           string `json:"username" binding:"required"`
	ArtworkName        string `json:"artworkName" binding:"required"`
	ArtworkDescription string `json:"artworkDescription"`
	ArtTag             []Tag  `json:"artTag"`
	ArtworkUrl         string `json:"artworkUrl" binding:"required"`
}

type UploadArtworkResponse struct {
	Log string `json:"log"`
}

type EditArtworkRequest struct {
	Username           string `json:"username" binding:"required"`
	ArtworkID          int    `json:"artworkID" binding:"required"`
	ArtworkName        string `json:"artworkName"`
	ArtworkDescription string `json:"artworkDescription"`
	ArtTag             []Tag  `json:"artTag"`
	ArtworkUrl         string `json:"artworkUrl"`
}

type DeleteArtworkRequest struct {
	Username  string `json:"username" binding:"required"`
	ArtworkID int    `json:"artworkID" binding:"required"`
}

type GetTagResponse struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	TagId   string `json:"tagID" binding:"required"`
	TagName string `json:"tagName" binding:"required"`
}

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

type SearchArtistRequest struct {
	ArtistName   string  `json:"artistName"`
	MinPriceRate int64   `json:"minPriceRate"`
	MaxPriceRate int64   `json:"maxPriceRate"`
	MinRating    float32 `json:"minRating"`
	MaxRating    float32 `json:"maxRating"`
	Tags         []Tag   `json:"tags"`
}

type SearchResultResponse struct {
	UserID   string  `json:"userID"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Rating   float32 `json:"rating"`
}

type ArtWorkResponse struct {
	ArtWorkID   string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UploadDate  string `json:"uploadDate"`
	Url         string `json:"url"`
	Tags        []Tag  `json:"tags"`
}

type ArtistProfileResponse struct {
	UserID          string            `json:"userID"`
	Username        string            `json:"username"`
	Name            string            `json:"name"`
	Surname         string            `json:"surname"`
	Rating          float32           `json:"rating"`
	MinPriceRate    int64             `json:"minPriceRate"`
	MaxPriceRate    int64             `json:"maxPriceRate"`
	Biography       string            `json:"biography"`
	ArtWorkResponse []ArtWorkResponse `json:"artworks"`
}
