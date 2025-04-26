package services

import (
	"strings"
	"sync"

	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
)

type SearchService struct {
	Store *repository.Store
}

func NewSearchService(store *repository.Store) *SearchService {
	return &SearchService{Store: store}
}

func (s *SearchService) Search(input string) models.SearchedData {
	var serched models.SearchedData
	input = strings.ToLower(input)
	input = strings.Trim(input, `"`) // ← this removes quotes from both sides
	s.Store.Wg.Add(3)
	go s.serchInArtists(input, &serched, &s.Store.Wg)
	go s.serchInLocations(input, &serched, &s.Store.Wg)
	go s.serchInmembers(input, &serched, &s.Store.Wg)
	s.Store.Wg.Wait()
	return serched
}

func (s *SearchService) serchInArtists(input string, searched *models.SearchedData, wg *sync.WaitGroup) {
	s.Store.Mutex.Lock()
	defer s.Store.Mutex.Unlock()
	defer wg.Done()
	for _, artist := range s.Store.Artists {
		to := strings.ToLower(artist.Name)
		if strings.Contains(to, input) {
			searched.Sugestions = append(searched.Sugestions, artist.Name)
			searched.Artists = append(searched.Artists, artist)
		}
	}
}

func (s *SearchService) serchInmembers(input string, searched *models.SearchedData, wg *sync.WaitGroup) {
	s.Store.Mutex.Lock()
	defer s.Store.Mutex.Unlock()
	defer wg.Done()
	for _, artist := range s.Store.Artists {
		is := false
		for _, mum := range artist.Members {
			to := strings.ToLower(mum)

			if strings.Contains(to, input) {
				searched.Sugestions = append(searched.Sugestions, mum)
				is = true
			}
		}
		if is {
			searched.Artists = append(searched.Artists, artist)
		}
	}
}

func (s *SearchService) serchInLocations(input string, searched *models.SearchedData, wg *sync.WaitGroup) {
	s.Store.Mutex.Lock()
	defer s.Store.Mutex.Unlock()
	defer wg.Done()
	for _, artist := range s.Store.Artists {
		is := false
		for _, loc := range s.Store.Locations.Index[artist.ID-1].Locations {
			to := strings.ToLower(loc)
			if strings.Contains(to, input) {
				searched.Sugestions = append(searched.Sugestions, loc)
				is = true
			}
		}
		if is {
			searched.Artists = append(searched.Artists, artist)
		}
	}
}
