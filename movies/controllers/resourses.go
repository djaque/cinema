package controllers

import (
	"github.com/djaque/cinema/movies/models"
)

type (
	// For Get - /movies
	MoviesResource struct {
		Data []models.Movie `json:"data"`
	}
	// For Post/Put - /movies
	MovieResource struct {
		Data models.Movie `json:"data"`
	}
)
