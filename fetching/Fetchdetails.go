package fetching

import (
	"fmt"
	"strconv"

	"groupie-tracker/models"
)

func Fetchdetails(Id string) (*models.Data, error) {
	var sarout models.Data
	if id, err := strconv.Atoi(Id); err != nil || id > 52 {
		return nil, fmt.Errorf("yguhhurfkhur")
	}
	artist, err := GetArtistById(Id)
	if err != nil {
		return nil, err
	}
	sarout.Id = artist.Id
	sarout.Image = artist.Image
	sarout.Name = artist.Name
	sarout.Members = artist.Members
	sarout.CreationDate = artist.CreationDate
	sarout.FirstAlbum = artist.FirstAlbum
	loc, err := Fetchlocal(Id)
	if err != nil {
		return nil, err
	}
	sarout.Loc = loc.Locations
	date, err := Fetchdates(Id)
	if err != nil {
		return nil, err
	}
	sarout.Date = date.Dates
	relation, err := Fetchrelations(Id)
	if err != nil {
		return nil, err
	}
	sarout.Datelocation = relation.Datelocation

	return &sarout, nil
}
