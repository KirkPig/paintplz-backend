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
	db.database.Raw("Call Register(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user_id,
		username,
		name,
		surname,
		email,
		citizenID,
		password,
		"artist",
		minPrice,
		maxPrice,
		biography,
	)
	return nil
}

func (db *GromDB) SeartArtist(name string, minPrice float64, maxPrice float64, minRate, maxRate float32, tag_name []string) (SearchArtistResponse, error) {
	query := `SELECT  A.ARTIST_USER_ID,
        U.NAME,
        U.SURNAME,
        A.RATING
		FROM    ARTIST A LEFT JOIN PAINTPLZ_USER U 
		ON U.PAINTPLZ_USER_ID = A.ARTIST_USER_ID
		WHERE     LOCATE(?, U.NAME) > 0 AND
        ? <= A.MIN_PRICE AND
        ? >= A.MAX_PRICE AND
        ? <= A.RATING AND
        ? >= A.RATING AND
        (? = "" OR EXISTS(
            SELECT W.ARTIST_USER_ID
            FROM ARTWORK W, ARTWORK_ARTTAG AA, ART_TAG T
            WHERE     A.ARTIST_USER_ID = W.ARTIST_USER_ID AND
                    W.ART_ID = AA.TAG_ID AND
                    T.TAG_ID = AA.TAG_ID AND
                    FIND_IN_SET(T.TAG_NAME, ?) > 0
        ))`
	tags := ""
	for i := 0; i < len(tag_name); i++ {
		tags += tag_name[i]
		if i+1 < len(tag_name) {
			tags += ","
		}
	}
	var result SearchArtistResponse
	err := db.database.Raw(query, name, minPrice, maxPrice, minRate, maxRate, tags).Scan(&result).Error
	return result, err
}

func (db *GromDB) RegisterCustomer(user_id, username, name, surname, email, citizenID, password string) error {

	db.database.Raw("Call Register(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",

		user_id,
		username,
		name,
		surname,
		email,
		citizenID,
		password,
		"customer",
		nil,
		nil,
		nil,
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
