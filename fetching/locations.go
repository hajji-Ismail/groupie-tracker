package fetching

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"groupie-tracker/models"
)

var local *[]models.Location

func init() {
	var err error
	local, err = fetchlocal("https://groupietrackers.herokuapp.com/api/locations")
	fmt.Println(local)
	
	if err != nil {

		log.Fatal(err)
	}

}

func fetchlocal(url string) (*[]models.Location, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		return nil, err
	}
	local, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var loc []models.Location
	err = json.Unmarshal(local, &loc)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &loc, nil
}
