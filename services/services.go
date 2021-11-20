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

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {
	response, err := s.database.Login(req.Username, req.Password)

	if response.UserType == "artist" {
		return LoginResponse{
			UserID:       response.UserID,
			Username:     response.Username,
			Name:         response.Name,
			Surname:      response.Surname,
			Email:        response.Email,
			CitizenID:    response.CitizenID,
			UserType:     true,
			MinPriceRate: response.MinPriceRate,
			MaxPriceRate: response.MaxPriceRate,
			Biography:    response.Biography,
		}, err
	}

	return LoginResponse{
		UserID:    response.UserID,
		Username:  response.Username,
		Name:      response.Name,
		Surname:   response.Surname,
		Email:     response.Email,
		CitizenID: response.CitizenID,
		UserType:  false,
	}, err
}

func (s *Service) ArtistProfile(userID string) (ArtistProfileResponse, error) {

	var profile ArtistProfileResponse

	user, err := s.database.GetArtistByID(userID)

	if err != nil {
		return profile, err
	}

	var artworks []ArtworkResponse

	artworks = make([]ArtworkResponse, 0)

	artworkTag, err := s.database.GetArtistArtwork(userID)

	if err != nil {
		return profile, err
	}

	var ma map[string][]Tag

	ma = make(map[string][]Tag)

	for _, art := range artworkTag {

		if art.TagID != "" {
			if _, ok := ma[art.ArtworkID]; ok {
				ma[art.ArtworkID] = append(ma[art.ArtworkID], Tag{
					TagId:   art.TagID,
					TagName: art.TagName,
				})
			} else {
				ma[art.ArtworkID] = append(make([]Tag, 0), Tag{
					TagId:   art.TagID,
					TagName: art.TagName,
				})
			}
		}

	}

	var m map[string]ArtworkResponse

	m = make(map[string]ArtworkResponse)

	for _, art := range artworkTag {
		if _, ok := m[art.ArtworkID]; !ok {
			var t []Tag

			if val, ok := ma[art.ArtworkID]; ok {
				t = val
			} else {
				t = make([]Tag, 0)
			}

			m[art.ArtworkID] = ArtworkResponse{
				ArtWorkID:   art.ArtworkID,
				Title:       art.Title,
				Description: art.Description,
				UploadDate:  art.UploadDate,
				Url:         art.ArtworkUrl,
				Tags:        t,
			}
		}
	}

	for _, val := range m {
		artworks = append(artworks, val)
	}

	profile = ArtistProfileResponse{
		UserID:          userID,
		Username:        user.Username,
		Name:            user.Name,
		Surname:         user.Surname,
		ProfileUrl:      user.ImageUrl,
		Rating:          0,
		MinPriceRate:    user.MinPriceRate,
		MaxPriceRate:    user.MaxPriceRate,
		Biography:       user.Biography,
		ArtWorkResponse: artworks,
	}

	return profile, nil

}

func (s *Service) GetAllTag() ([]Tag, error) {
	response, err := s.database.GetAllTag()
	var t []Tag
	t = make([]Tag, 0)

	if err != nil {
		return t, err
	}

	for _, val := range response {
		t = append(t, Tag{
			TagId:   val.TagID,
			TagName: val.TagName,
		})
	}

	return t, nil

}
