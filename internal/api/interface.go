package api

import "github.com/jolinGalal/andela/internal/model"

// API ...
type API struct {
}

func NewAPI() APIInterface {
	return &API{}
}

// APIInterface ...
type APIInterface interface {
	GetPosts() (*[]model.Post, error)
	GetComments() (*[]model.Comment, error)
}
