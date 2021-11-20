package services

import (
	"github.com/KirkPig/paintplz-backend/repository"
	uuid "github.com/nu7hatch/gouuid"
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
	new_uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	if req.UserType {
		err = s.database.RegisterArtist(new_uuid.String(), req.Username, req.Name, req.Surname, req.Email, req.CitizenID, req.Password,
			req.MinPriceRate, req.MaxPriceRate, req.Biography)
	} else {
		err = s.database.RegisterCustomer(new_uuid.String(), req.Username, req.Name, req.Surname, req.Email, req.CitizenID, req.Password)
	}

	return err

}

func (s *Service) Login(req LoginRequest) (repository.LoginDBResponse, error) {
	response, err := s.database.Login(req.Username, req.Password)
	return response, err
}
