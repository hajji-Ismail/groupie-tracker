package models

type Data struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Loc          []string
	Date         []string
	Datelocation map[string][]string 
}


type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locationss"`

	ConcertDates string `json:"concertDates"`
	Relations    string `json:"relations"`
}
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type  Relation struct {
	ID           int                 `json:"id"`
	Datelocation map[string][]string `json:"datesLocations"`
}
