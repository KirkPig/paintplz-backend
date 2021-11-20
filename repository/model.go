package repository

type RegisterDBResponse struct {
	UserID    string `gorm:"column:PAINTPLZ_USER_ID"`
	Username  string `gorm:"column:USERNAME"`
	Name      string `gorm:"column:NAME"`
	Surname   string `gorm:"column:SURNAME"`
	Email     string `gorm:"column:EMAIL"`
	CitizenID string `gorm:"column:CITIZEN_ID"`
	ImageUrl  string `gorm:"column:PROFILE_URL"`
	UserType  string `gorm:"column:USER_TYPE"`
}

type LoginDBResponse struct {
	UserID       string  `gorm:"column:PAINTPLZ_USER_ID"`
	Username     string  `gorm:"column:USERNAME"`
	Name         string  `gorm:"column:NAME"`
	Surname      string  `gorm:"column:SURNAME"`
	Email        string  `gorm:"column:EMAIL"`
	CitizenID    string  `gorm:"column:CITIZEN_ID"`
	ImageUrl     string  `gorm:"column:PROFILE_URL"`
	UserType     string  `gorm:"column:USER_TYPE"`
	MinPriceRate float64 `gorm:"column:MIN_PRICE"`
	MaxPriceRate float64 `gorm:"column:MAX_PRICE"`
	Biography    string  `gorm:"column:BIOGRAPHY"`
}

type Artist struct {
	UserID       string
	Username     string
	Name         string
	Surname      string
	Email        string
	CitizenID    string
	ImageUrl     string
	UserType     string
	MinPriceRate float64
	MaxPriceRate float64
	Biography    string
}

type ArtworkDB struct {
	ArtworkID   string
	Title       string
	Description string
	UploadDate  string
	ArtworkUrl  string
	TagID       string
	TagName     string
}
