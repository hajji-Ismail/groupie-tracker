package fetching

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"groupie-tracker/models"
)

var Locations *[]models.Locations

func init() {
	var err  error
	Locations, err = fetchinlocation("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	
}

func fetchinlocation(url string) (*[]models.Locations, error) { // Change return type to a slice
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err // Return nil for the slice on error
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		return nil, err // Return nil for the slice on non-200 response
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err // Return nil for the slice on error
	}

	var location []models.Locations // Change to a slice of Artist
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Println(err)
		return nil, err // Return nil for the slice on error
	}

	return &location, nil // Return the slice of artists
}