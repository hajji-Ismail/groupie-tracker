package fetching

import (
	"groupie-tracker/models"
	"log"
	"strconv"
)

func GetArtistById(idstr string) *models.Artist {

	var artist *models.Artist
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println("invalid Id")
	}

	artists := Artists

	// Dereference the pointer to access the slice
	for _, a := range *artists {
		if a.Id == id { // Assuming each artist has an ID field
			artist = &a
			break
		}
	}

	if artist == nil {
		log.Fatal("there is no artist by the id you are giving")
		return nil
	}

	return artist
}
