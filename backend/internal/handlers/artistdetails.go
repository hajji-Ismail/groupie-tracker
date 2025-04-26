package handlers

import (
	"net/http"
	"strconv"

	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
)

type ArtistsDetailsHandler struct {
	Service *services.ArtistsDetailsService
}

func NewArtistDetailsService(service *services.ArtistsDetailsService) *ArtistsDetailsHandler {
	return &ArtistsDetailsHandler{Service: service}
}

func (s *ArtistsDetailsHandler) GetArtistDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "")
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Bad Request", "invalid artist id")
		return
	}
	details, err := s.Service.GetArtistDetails(id)
	if err != nil {
		// logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Bad Request", "artist my be deleted")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, details)
}
