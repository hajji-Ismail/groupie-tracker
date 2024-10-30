package fetch

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"groupie-tracker/models"
)

// fetchAPIData fetches data from the specified API URL and returns it
type Data struct {
    Artists *[]models.Artist
}
var Fetch  Data
func FetchAPIData()  {
	

	// Create a new HTTP request
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: received status code %d", resp.StatusCode)
		
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		
	}

	// Unmarshal the JSON response into the struct
	err = json.Unmarshal(body, &Fetch.Artists)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
	
	}


}