package models

type FilterRequest struct {
	CreationFrom string `json:"creationDateFrom"`
	CreationTo   string `json:"creationDateTo"`
	AlbumFrom    string `json:"firstAlbumFrom"`
	AlbumTo      string `json:"firstAlbumTo"`
	Members      string `json:"members"`
	ConcertDate  string `json:"concertDates"`
}
// filed to filter
