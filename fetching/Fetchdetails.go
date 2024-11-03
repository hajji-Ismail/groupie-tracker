package fetching

import "groupie-tracker/models"

func Fetchdetails( Id string) (*models.Data, error) {
	var sarout models.Data
	
	artist := GetArtistById(Id)
	sarout.Id = artist.Id
	sarout.Image = artist.Image
	sarout.Name = artist.Name
	sarout.Members = artist.Members
	sarout.CreationDate =  artist.CreationDate
	sarout.FirstAlbum = artist.FirstAlbum
	loc,err := Fetchlocal(Id)

	if err != nil{
		return nil, err
	}
	sarout.Loc = loc.Locations
	return &sarout, nil



}