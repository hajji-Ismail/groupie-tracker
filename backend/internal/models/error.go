package models

type APIError struct {
	Status  int
	Message string
	Details string
}
