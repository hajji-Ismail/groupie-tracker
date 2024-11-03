package fetching

import (
	"encoding/json"
	"groupie-tracker/models"
	"io"
	"log"
	"net/http"
)

var Artists *[]models.Artist

func init() {
	var err error
	Artists, err = fetching("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}

}

func fetching(url string) (*[]models.Artist, error) { // Change return type to a slice
	resp, err := http.Get(url)
	if err != nil {

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

	var artists []models.Artist // Change to a slice of Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Println(err)
		return nil, err // Return nil for the slice on error
	}

	return &artists, nil // Return the slice of artists
}
