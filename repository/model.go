package repository

type RegisterDBResponse struct {
	UserID    string `json:"userID"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	CitizenID string `json:"citizenID"`
	ImageUrl  string `json:"imageUrl"`
	UserType  string `json:"userType"`
}

type LoginDBResponse struct {
	UserID       string  `json:"userID"`
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	CitizenID    string  `json:"citizenID"`
	ImageUrl     string  `json:"imageUrl"`
	UserType     string  `json:"userType"`
	MinPriceRate float64 `json:"minPriceRate"`
	MaxPriceRate float64 `json:"maxPriceRate"`
	Biography    string  `json:"biography"`
}

type SearchArtistResponse struct {
	UserID   string  `json:"userID"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Rating   float32 `json:"rating"`
}
