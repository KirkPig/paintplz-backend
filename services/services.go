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

func (s *Service) ArtistProfile(userID string) (ArtistProfileResponse, error) {

	var profile ArtistProfileResponse

	user, err := s.database.GetArtistByID(userID)

	if err != nil {
		return profile, err
	}

	var artworks []ArtworkResponse

	_, err = s.database.GetArtistArtwork(userID)

	if err != nil {
		return profile, err
	}

	profile = ArtistProfileResponse{
		UserID:          user.UserID,
		Username:        user.Username,
		Name:            user.Name,
		Surname:         user.Surname,
		Rating:          0,
		MinPriceRate:    user.MinPriceRate,
		MaxPriceRate:    user.MaxPriceRate,
		Biography:       user.Biography,
		ArtWorkResponse: artworks,
	}

	return profile, nil

}
