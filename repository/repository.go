package repository

import (
	"log"

	"github.com/jinzhu/gorm"
)

type GromDB struct {
	database *gorm.DB
}

func New(db *gorm.DB) *GromDB {
	return &GromDB{database: db}
}

func (db *GromDB) RegisterArtist(user_id, username, name, surname, email, citizenID, password string, minPrice, maxPrice float64, biography string) error {
	var newUser RegisterDBResponse
	err := db.database.Raw("Call Register(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
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
	).Scan(&newUser).Error
	return err
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
	var newUser RegisterDBResponse
	err := db.database.Raw("Call Register(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",

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
	).Scan(&newUser).Error

	return err
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

	log.Println(loginQuery)

	return loginQuery, err
}

func (db *GromDB) GetArtistByID(userID string) (Artist, error) {

	var artist Artist

	query := `SELECT	U.USERNAME, 
	U.NAME, 
	U.SURNAME, 
	U.EMAIL, 
	U.CITIZEN_ID,
	U.PROFILE_URL,
	U.USER_TYPE, 
	A.MIN_PRICE, 
	A.MAX_PRICE, 
	A.BIOGRAPHY
FROM PAINTPLZ_USER U, ARTIST A
WHERE U.PAINTPLZ_USER_ID = ? AND A.ARTIST_USER_ID = ?;`

	err := db.database.Raw(query, userID, userID).Scan(&artist).Error

	return artist, err

}

func (db *GromDB) GetArtistArtwork(userID string) ([]ArtworkDB, error) {

	var artwork []ArtworkDB

	query := `SELECT  W.ART_ID,
	W.ART_TITLE,
	W.ART_DESC,
	W.UPLOAD_DATE,
	W.ART_URL,
	T.TAG_ID,
	T.TAG_NAME
FROM ARTWORK W, ART_TAG T, ARTWORK_ARTTAG AA
WHERE W.ARTIST_USER_ID = ? AND W.ART_ID = AA.ART_ID AND T.TAG_ID = AA.TAG_ID
ORDER BY W.ART_ID DESC, T.TAG_ID;`

	err := db.database.Raw(query, userID).Scan(&artwork).Error

	return artwork, err

}

func (db *GromDB) GetAllTag() ([]Tag, error) {

	var tags []Tag
	query := `SELECT TAG_ID, TAG_NAME
	FROM ART_TAG;`

	err := db.database.Raw(query).Scan(&tags).Error

	return tags, err

}

func (db *GromDB) UploadArtwork(userID, artID, artTitle, artDesc, url string, tag_id, tag_name []string) error {
	query := `INSERT INTO ARTWORK(ART_ID,ART_URL,ART_TITLE,ART_DESC,ARTIST_USER_ID)
	VALUES (?, ?, ?, ?, ?)`
	/**
	query := `INSERT INTO ARTWORK(ART_ID,ART_URL,ART_TITLE,ART_DESC,ARTIST_USER_ID)
	VALUES (@artId, @artUrl, @artTitle, @artDesc, @userID)`
	*/
	uploadErr := db.database.Raw(query, artID, url, artTitle, artDesc, userID).Error
	if uploadErr != nil {
		return uploadErr
	}
	for i := 0; i < len(tag_id); i += 1 {
		addTagErr := db.database.Raw("call ADD_TAG(?, ?, ?)", artID, tag_id[i], tag_name[i]).Error
		if addTagErr != nil {
			return addTagErr
		}
	}
	return nil
}

func (db *GromDB) EditArtwork(userID, artID, artName, artDesc, url string, tag_id, tag_name []string) error {
	/**
		CREATE PROCEDURE ART_EDIT (
		IN ARTWORKID VARCHAR(100),
		IN ARTWORKNAME VARCHAR(100),
		IN ARTWORKDESCRIPTION TEXT(500),
		IN ARTWORKURL VARCHAR(100)
	)	*/
	editErr := db.database.Raw("call ART_EDIT(?, ?, ?, ?)", artID, artName, artDesc, url).Error
	if editErr != nil {
		return editErr
	}
	/**
	CREATE PROCEDURE ART_TAG_EDIT (
	IN ARTWORKID VARCHAR(100),
	IN ARTWORKTAGID VARCHAR(100)
	IN ARTWORKTAGNAME VARCHAR(100)
	*/
	for i := 0; i < len(tag_id); i += 1 {
		editTagErr := db.database.Raw("call ART_TAG_EDIT(?, ?, ?)", artID, tag_id[i], tag_name[i]).Error
		if editTagErr != nil {
			return editTagErr
		}
	}
	return nil
}

func (db *GromDB) DeleteArtwork(artworkID, artistUserID string) error {
	query := `DELETE FROM ARTWORK where Art_id = ? and  artistUserID = ?`

	return db.database.Raw(query, artworkID, artistUserID).Error
}
