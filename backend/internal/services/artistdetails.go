package services

import (
	"errors"
	"fmt"

	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
)

type ArtistsDetailsService struct {
	Store *repository.Store
}

func NewArtistDetailsService(store *repository.Store) *ArtistsDetailsService {
	return &ArtistsDetailsService{Store: store}
}

func (s *ArtistsDetailsService) GetArtistDetails(id int) (models.ArtistDetails, error) {
	if id < 1 || id > 52 {
		//utils.RespondWithError(w,http.StatusBadRequest,"Bad Request","")
		return models.ArtistDetails{}, fmt.Errorf("bad request")
	}
	var details models.ArtistDetails
	var err error
	s.Store.Wg.Add(4)

	go func() {
		defer s.Store.Wg.Done()
		artist, ok := s.getArtistByID(id)
		if ok {
			details.ArtistInfo = artist
		} else {
			s.Store.Mutex.Lock()
			err = errors.New("artist not found")
			s.Store.Mutex.Unlock()
		}
	}()

	go func() {
		defer s.Store.Wg.Done()
		location, ok := s.getLocationById(id)
		if ok {
			details.Locations = location
		} else {
			s.Store.Mutex.Lock()
			err = errors.New("location not found")
			s.Store.Mutex.Unlock()
		}
	}()

	go func() {
		defer s.Store.Wg.Done()
		date, ok := s.getDateById(id)
		if ok {
			details.Dates = date
		} else {
			s.Store.Mutex.Lock()
			err = errors.New("date not found")
			s.Store.Mutex.Unlock()
		}
	}()

	go func() {
		defer s.Store.Wg.Done()
		relation, ok := s.getRealtionById(id)
		if ok {
			details.Relations = relation
		} else {
			s.Store.Mutex.Lock()
			err = errors.New("relation not found")
			s.Store.Mutex.Unlock()
		}
	}()

	s.Store.Wg.Wait()

	if err != nil {
		return models.ArtistDetails{}, err
	}

	return details, nil
}

func (s *ArtistsDetailsService) getArtistByID(id int) (models.Artist, bool) {
	s.Store.Mutex.Lock()
	artist := s.Store.Artists[id-1]
	s.Store.Mutex.Unlock()
	return artist, true
}

func (s *ArtistsDetailsService) getLocationById(id int) (models.Location, bool) {
	s.Store.Mutex.Lock()
	location := s.Store.Locations.Index[id-1]
	s.Store.Mutex.Unlock()
	return location, true
}

func (s *ArtistsDetailsService) getDateById(id int) (models.Date, bool) {
	s.Store.Mutex.Lock()
	Date := s.Store.Dates.Index[id-1]
	s.Store.Mutex.Unlock()
	return Date, true
}

func (s *ArtistsDetailsService) getRealtionById(id int) (models.Relation, bool) {
	s.Store.Mutex.Lock()
	Relations := s.Store.Realtions.Index[id-1]
	s.Store.Mutex.Unlock()
	return Relations, true
}
