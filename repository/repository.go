package repository

import (
	"github.com/jinzhu/gorm"
)

type GromDB struct {
	database *gorm.DB
}

func New(db *gorm.DB) *GromDB {
	return &GromDB{database: db}
}

func (db *GromDB) RegisterArtist(user_id, username, name, surname, email, citizenID, password string, minPrice, maxPrice float64, biography string) error {
	db.database.Raw("Call Register(@user_id, @username, @name, @surname, @email, @citizenID, @password, @usertype, @minPrice, @maxPrice, @biography)",
		map[string]interface{}{
			"user_id":   user_id,
			"username":  username,
			"name":      name,
			"surname":   surname,
			"email":     email,
			"citizenID": citizenID,
			"password":  password,
			"usertype":  "artist",
			"minPrice":  minPrice,
			"maxPrice":  maxPrice,
			"biography": biography,
		},
	)
	return nil
}

func (db *GromDB) RegisterCustomer(user_id, username, name, surname, email, citizenID, password string) error {

	db.database.Raw("Call Register(@user_id, @username, @name, @surname, @email, @citizenID, @password, @usertype, @minPrice, @maxPrice, @biography)",
		map[string]interface{}{
			"user_id":   user_id,
			"username":  username,
			"name":      name,
			"surname":   surname,
			"email":     email,
			"citizenID": citizenID,
			"password":  password,
			"usertype":  "customer",
			"minPrice":  nil,
			"maxPrice":  nil,
			"biography": nil,
		},
	)

	return nil
}

func (db *GromDB) Login(username, password string) (LoginDBResponse, error) {

	var loginQuery LoginDBResponse
	query := `SELECT	U.PAINTPLZ_USER_ID, 
	U.USERNAME, 
    U.NAME, 
	U.SURNAME, 
	U.EMAIL, 
	U.CITIZEN_ID,
	U.PROFILE_URL,
	U.USER_TYPE, 
	A.MIN_PRICE, 
	A.MAX_PRICE, 
	A.BIOGRAPHY
FROM PAINTPLZ_USER U LEFT JOIN ARTIST A
ON U.PAINTPLZ_USER_ID = A.ARTIST_USER_ID
WHERE U.USERNAME = ? AND U.PASSWORD = ?`

	err := db.database.Raw(query, username, password).Scan(&loginQuery).Error

	return loginQuery, err
}
