package api

import (
	"net/http"

	"go-rest-api/internal/dependencies"
)

func NewRouter(deps *dependencies.Dependencies) http.Handler {
	mux := http.NewServeMux()

	// Register routes by domain
	registerArtistsRoutes(mux, deps)
	// registerUserRoutes(mux, deps)

	return mux
}

func registerArtistsRoutes(mux *http.ServeMux, deps *dependencies.Dependencies) {
	mux.HandleFunc("/api/artists", deps.AllArtistsHandler.GetAllArtists)
	mux.HandleFunc("/api/artists/filter", deps.FilterHandler.Filter)
	mux.HandleFunc("/api/artists/details", deps.ArtistDetailsHandler.GetArtistDetails)
	mux.HandleFunc("/api/artists/search", deps.SearchHandler.Search)
}
