package services

type UploadArtworkRequest struct {
	UserID             string `json:"userID" binding:"required"`
	ArtworkName        string `json:"artworkName" binding:"required"`
	ArtworkDescription string `json:"artworkDescription"`
	ArtTag             []Tag  `json:"artTag"`
	ArtworkUrl         string `json:"artworkUrl" binding:"required"`
}

type EditArtworkRequest struct {
	UserID             string `json:"userID" binding:"required"`
	ArtworkID          string `json:"artworkID" binding:"required"`
	ArtworkName        string `json:"artworkName"`
	ArtworkDescription string `json:"artworkDescription"`
	ArtTag             []Tag  `json:"artTag"`
	ArtworkUrl         string `json:"artworkUrl"`
}

type DeleteArtworkRequest struct {
	UserID    string `json:"userID" binding:"required"`
	ArtworkID string `json:"artworkID" binding:"required"`
}

type GetTagResponse struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	TagId   string `json:"tagID" binding:"required"`
	TagName string `json:"tagName" binding:"required"`
}

type RegisterRequest struct {
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	CitizenID    string  `json:"citizenID"`
	Password     string  `json:"password"`
	UserType     bool    `json:"userType"`
	MinPriceRate float64 `json:"minimumPriceRate,omitempty"`
	MaxPriceRate float64 `json:"maximumPriceRate,omitempty"`
	Biography    string  `json:"biography,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID       string  `json:"userID" binding:"required"`
	Username     string  `json:"username" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Surname      string  `json:"surname" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	CitizenID    string  `json:"citizenID" binding:"required"`
	UserType     bool    `json:"userType" binding:"required"`
	MinPriceRate float64 `json:"minPriceRate,omitempty"`
	MaxPriceRate float64 `json:"maxPriceRate,omitempty"`
	Biography    string  `json:"biography,omitempty"`
}

type SearchArtistRequest struct {
	ArtistName   string   `json:"artistName"`
	MinPriceRate *int64   `json:"minimumPriceRate"`
	MaxPriceRate *int64   `json:"maximumPriceRate"`
	MinRating    *float32 `json:"minimumRating"`
	MaxRating    *float32 `json:"maximumRating"`
	Tags         []Tag    `json:"tags"`
}

type SearchResult struct {
	UserID     string  `json:"userID"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Rating     float32 `json:"rating"`
	ProfileUrl string  `json:"profileUrl"`
}

type SearchResultResponse struct {
	SearchResult []SearchResult `json:"searchResult"`
}

type ArtworkResponse struct {
	ArtWorkID   string `json:"artworkID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UploadDate  string `json:"uploadDate"`
	Url         string `json:"url"`
	Tags        []Tag  `json:"artTag"`
}

type ArtistProfileResponse struct {
	UserID               string            `json:"userID"`
	Username             string            `json:"username"`
	Name                 string            `json:"name"`
	Surname              string            `json:"surname"`
	ProfileUrl           string            `json:"profileUrl"`
	Rating               float32           `json:"rating"`
	MinPriceRate         float64           `json:"minimumPriceRate"`
	MaxPriceRate         float64           `json:"maximumPriceRate"`
	Biography            string            `json:"biography"`
	ArtWorkResponse      []ArtworkResponse `json:"artworks"`
	MongoArtworkResponse []ArtworkResponse `json:"mongoArtworks"`
}

type Logger struct {
	Log string `json:"log"`
}
