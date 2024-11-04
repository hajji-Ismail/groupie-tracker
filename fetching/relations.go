package fetching

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"groupie-tracker/models"
)


func Fetchrelations(id string) (*models.Relation, error) {
	url := "https://groupietrackers.herokuapp.com/api/relation" + "/" + id
	resp, err := http.Get(url)
	
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		return nil, fmt.Errorf("failed to fetch relations: %s", resp.Status)
	}
	
	relationData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rel models.Relation
	err = json.Unmarshal(relationData, &rel)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	return &rel, nil
}