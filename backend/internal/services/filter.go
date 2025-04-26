package services

import (
	"fmt"
	"strconv"
	"strings"

	"go-rest-api/internal/models"
	"go-rest-api/internal/repository"
)

type FilterService struct {
	Store *repository.Store
}

func NewFilterService(store *repository.Store) *FilterService {
	return &FilterService{Store: store}
}

func (s *FilterService) Filter(data models.FilterRequest) ([]models.Artist, error) {
	var filteredArtists []models.Artist
	// Parse filter values
	creationFrom, err1 := strconv.Atoi(data.CreationFrom)
	creationTo, err2 := strconv.Atoi(data.CreationTo)
	albumFrom, err3 := strconv.Atoi(data.AlbumFrom)
	albumTo, err4 := strconv.Atoi(data.AlbumTo)
	MembersArr := []int{}
	for _, member := range data.Members {
		mem, err := strconv.Atoi(member)
		if err != nil {
			return nil, fmt.Errorf("error parsing number atoi error")
		}
		MembersArr = append(MembersArr, mem)

	}

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, fmt.Errorf("error parsing number atoi error")
	}

	for _, artist := range s.Store.Artists {
		isMatch := true
		// Creation date filter
		if creationFrom != 0 || creationTo != 0 {
			if creationFrom != 0 && artist.CreationDate < creationFrom {
				isMatch = false
			}
			if creationTo != 0 && artist.CreationDate > creationTo {
				isMatch = false
			}
		}

		// First album year filter
		if albumFrom != 0 || albumTo != 0 {
			parts := strings.Split(artist.FirstAlbum, "-")
			if len(parts) != 3 {
				isMatch = false
				continue
			}

			albumYear, err := strconv.Atoi(parts[2])
			if err != nil {
				isMatch = false
				continue
			}

			if albumFrom != 0 && albumYear < albumFrom {
				isMatch = false
			}
			if albumTo != 0 && albumYear > albumTo {
				isMatch = false
			}
		}

		// Members filter

		// Members filter
		if len(MembersArr) != 0 {
			matchFound := false
			for _, val := range MembersArr {
				if len(artist.Members) == val {
					matchFound = true
					break
				}
			}
			if !matchFound {
				isMatch = false
			}
		}

		// Concert location filter
		if data.ConcertDate != "" {
			if !s.matchLocation(artist.ID, data.ConcertDate) {
				isMatch = false
			}
		}

		if isMatch {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	return filteredArtists, nil
}

func (f *FilterService) matchLocation(id int, query string) bool {
	query = strings.ToLower(strings.TrimSpace(query))
	query = strings.ReplaceAll(query, ",", "-")
	query = strings.ReplaceAll(query, " ", "")

	if id <= 0 || id > len(f.Store.Locations.Index) {
		return false
	}

	for _, loc := range f.Store.Locations.Index[id-1].Locations {
		normalizedLoc := strings.ToLower(strings.TrimSpace(loc))
		normalizedLoc = strings.ReplaceAll(normalizedLoc, ",", "-")
		if strings.Contains(normalizedLoc, query) {
			return true
		}
	}
	return false
}
