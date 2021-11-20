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

type ArtworkDB struct {
	ArtworkID   string `gorm:"column:ART_ID"`
	Title       string `gorm:"column:ART_TITLE"`
	Description string `gorm:"column:ART_DESC"`
	UploadDate  string `gorm:"column:UPLOAD_DATE"`
	ArtworkUrl  string `gorm:"column:ART_URL"`
	TagID       string `gorm:"column:TAG_ID"`
	TagName     string `gorm:"column:TAG_NAME"`
}

type Tag struct {
	TagID   string `gorm:"column:TAG_ID"`
	TagName string `gorm:"column:TAG_NAME"`
}

type SearchArtistResponse struct {
	UserID   string  `json:"userID"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Rating   float32 `json:"rating"`
}
