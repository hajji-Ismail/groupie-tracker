package services

import (
	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
)

type AllArtistsService struct {
	Store *repository.Store
}

func NewAllArtistsService(store *repository.Store) *AllArtistsService {
	return &AllArtistsService{Store: store}
}

func (s *AllArtistsService) GetAllArtists() []models.Artist {
	s.Store.Mutex.Lock()
	defer s.Store.Mutex.Unlock()
	return s.Store.Artists
}
