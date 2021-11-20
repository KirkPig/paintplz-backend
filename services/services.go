package services

import (
	"github.com/KirkPig/paintplz-backend/repository"
)

type Service struct {
	database repository.GromDB
}

func NewService(db repository.GromDB) *Service {
	return &Service{
		database: db,
	}
}

func (s *Service) Register(req RegisterRequest) error {

	var err error

	if req.UserType {
		err = s.database.RegisterArtist()
	} else {
		err = s.database.RegisterCustomer()
	}

	return err

}
