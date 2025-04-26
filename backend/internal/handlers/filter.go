package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-rest-api/internal/models"
	"go-rest-api/internal/services"
	"go-rest-api/internal/utils"
)

type FilterHandler struct {
	Service *services.FilterService
}

func NewFilterHandler(filterservice *services.FilterService) *FilterHandler {
	return &FilterHandler{Service: filterservice}
}

func (h *FilterHandler) Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "")
		return
	}
	fmt.Println(r.Method)
	var data models.FilterRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		// logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Bad Request", "")
		return
	}
	artists, err := h.Service.Filter(data)
	if err != nil {
		// logger.LogWithDetails(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Bad Request", "")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, artists)
}
