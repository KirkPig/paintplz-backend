package repository

import (
	"errors"

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
	U.USER_TYPE, 
	A.MIN_PRICE, 
	A.MAX_PRICE, 
	A.BIOGRAPHY
FROM PAINTPLZ_USER U LEFT JOIN ARTIST A
WHERE U.USERNAME = @Username AND U.PASSWORD = @Password
ON U.PAINTPLZ_USER_ID = A.ARTIST_USER_ID"`

	err := db.database.Raw(query,
		map[string]interface{}{
			"Username": username,
			"Password": password,
		}).Scan(&loginQuery).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return loginQuery, nil
	} else {
		return loginQuery, err
	}
}
