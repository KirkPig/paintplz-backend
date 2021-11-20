package repository

type LoginDBResponse struct {
	UserID       string  `json:"userID"`
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	CitizenID    string  `json:"citizenID"`
	UserType     bool    `json:"userType"`
	MinPriceRate float64 `json:"minPriceRate"`
	MaxPriceRate float64 `json:"maxPriceRate"`
	Biography    string  `json:"biography"`
}
