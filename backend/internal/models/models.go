package models

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	FirstAlbum   string   `json:"firstAlbum"`
	CreationDate int      `json:"creationDate"`
	// Locations     string `json:"locations"`    // api
	// ConcertDates  string `json:"concertDates"` // api
	// Relations     string `json:"relations"`    // api
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationsApiRespnse struct {
	Index []Location `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type DatesApiResponse struct {
	Index []Date `json:"index"`
}

type Relation struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type RelationsApiResponse struct {
	Index []Relation `json:"index"`
}


