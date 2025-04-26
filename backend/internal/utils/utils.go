package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"go-rest-api/internal/models"
	"go-rest-api/pkg/logger"
)

// FetchJSON fetches JSON from a given URL
func Fetch(url string, target any, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		logger.LogWithDetails(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.LogWithDetails(fmt.Errorf("failed to fetch data from %s: StatusCode %d", url, resp.StatusCode))
		return
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogWithDetails(err)
		return
	}
	// Parse the JSON response into the appropriate struct
	if err := json.Unmarshal(body, &target); err != nil {
		logger.LogWithDetails(err)
		return
	}
}

// Helper function to standardize error responses
func RespondWithError(w http.ResponseWriter, statusCode int, message, details string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	var err = models.APIError{
		Status:  statusCode,
		Message: message,
		Details: details,
	}
	if err := json.NewEncoder(w).Encode(err); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Errror", "")
	}
}

// Helper function to standardize JSON responses
func RespondWithJSON(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.LogWithDetails(err)
		RespondWithError(w, http.StatusInternalServerError, "internal server error", "")
		
	}
}
