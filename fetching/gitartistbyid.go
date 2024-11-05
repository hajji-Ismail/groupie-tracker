package fetching

import (
	"strconv"

	"groupie-tracker/models"
)

func GetArtistById(idstr string) (*models.Artist, error) {
	var artist *models.Artist
	id, err := strconv.Atoi(idstr)

	if err != nil || id > 52 {
		return nil, err
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
		return nil, err
	}

	return artist, nil
}
