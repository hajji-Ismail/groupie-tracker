package repository

import (
	"fmt"
	"sync"

	"go-rest-api/internal/models"
	"go-rest-api/internal/utils"
)

type Store struct {
	Artists   []models.Artist
	Locations models.LocationsApiRespnse
	Realtions models.RelationsApiResponse
	Dates     models.DatesApiResponse
	Wg        sync.WaitGroup
	Mutex     sync.Mutex
}

func New_Store() *Store {
	return &Store{}
}

// loadd data
func (s *Store) LoadData() {
	apiUrls := []string{
		"https://groupietrackers.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/locations",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relation",
	}
	s.Wg.Add(len(apiUrls))
	go utils.Fetch(apiUrls[0], &s.Artists, &s.Wg)
	go utils.Fetch(apiUrls[1], &s.Locations, &s.Wg)
	go utils.Fetch(apiUrls[2], &s.Dates, &s.Wg)
	go utils.Fetch(apiUrls[3], &s.Realtions, &s.Wg)
	s.Wg.Wait()
	// fmt.Println("Fetched Artists Data:", s.Artists[0])
	// fmt.Println("Fetched Locations Data:", s.Locations.Index[0])
	// fmt.Println("Fetched Dates Data:", s.Dates.Index)
	// fmt.Println("Fetched Relation Data:", s.Realtions.Index)
	fmt.Println("seccefully loaded data")
}
