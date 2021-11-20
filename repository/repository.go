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

func (db *GromDB) Register() error {
	return nil
}

func (db *GromDB) Login(username, password string) (LoginDBResponse, error) {

	var loginQuery LoginDBResponse

	err := db.database.Raw("", username, password).Scan(&loginQuery).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return loginQuery, nil
	} else {
		return loginQuery, err
	}
}
