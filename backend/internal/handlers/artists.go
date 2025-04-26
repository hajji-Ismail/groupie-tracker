package handlers

import (
	"net/http"

	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
)

type AllArtistsHandler struct {
	Service *services.AllArtistsService
}

func NewAllArtistsHandler(allartistservice *services.AllArtistsService) *AllArtistsHandler {
	return &AllArtistsHandler{Service: allartistservice}
}

func (h *AllArtistsHandler) GetAllArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "")
		return
	}
	artists := h.Service.GetAllArtists()
	utils.RespondWithJSON(w, http.StatusOK, artists)
}
