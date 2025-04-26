package handlers

import (
	"net/http"

	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
)

type SearchHandler struct {
	Service *services.SearchService
}

func NewSearchHandler(filterservice *services.SearchService) *SearchHandler {
	return &SearchHandler{Service: filterservice}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "")
		return
	}
	input := r.URL.Query().Get("find")
	if input == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Bad Request", "")
		return
	}
	data := h.Service.Search(input)
	utils.RespondWithJSON(w, http.StatusOK, data)
}
