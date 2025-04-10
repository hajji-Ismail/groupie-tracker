package fetching

import (
	"encoding/json"
	
	"io"
	"log"
	"net/http"

	"groupie-tracker/models"
)



func Fetchlocal(id string) (*models.Location , error)   {
	url := "https://groupietrackers.herokuapp.com/api/locations" + "/" + id
	resp, err := http.Get(url)

	if err != nil {
		return nil,err
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


	var loc models.Location
	err = json.Unmarshal(local, &loc)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &loc, nil
}
