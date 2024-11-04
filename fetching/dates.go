package fetching

import (
	"encoding/json"
	
	"io"
	"log"
	"net/http"

	"groupie-tracker/models"
)



func Fetchdates(id string) (*models.Date , error)   {
	url := "https://groupietrackers.herokuapp.com/api/dates" + "/" + id
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
	

	var Date models.Date
	err = json.Unmarshal(local, &Date)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Date, nil
}
