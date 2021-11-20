package services

import "github.com/KirkPig/paintplz-backend/repository"

type Service struct {
	database repository.GromDB
}

func NewService(db repository.GromDB) *Service {
	return &Service{
		database: db,
	}
}

func (s *Service) Register(req RegisterRequest) error {

	err := s.database.Register()

	if err != nil {

	}

	return nil

}
