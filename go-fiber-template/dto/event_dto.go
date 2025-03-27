package dto

import (
	"github.com/tapeds/go-fiber-template/entity"
)

type (
	EventCreateRequest struct {
		Name        string `json:"name"`
		AuthorID    string `json:"author_id"`
		Price       int    `json:"price"`
		Capacity    int    `json:"capacity"`
		Availabilty int    `json:"availabilty"`
	}

	EventResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		AuthorID    string `json:"author_id"`
		AuthorName  string `json:"author_name"`
		Price       int    `json:"price"`
		Capacity    int    `json:"capacity"`
		Availabilty int    `json:"availabilty"`
	}

	EventPaginationResponse struct {
		Data []EventResponse `json:"data"`
		PaginationResponse
	}

	GetAllEventRepositoryResponse struct {
		Events []entity.Event
		PaginationResponse
	}

	EventUpdateRequest struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Capacity    int    `json:"capacity"`
		Availabilty int    `json:"availabilty"`
	}

	EventByIdRequest struct {
		ID string `json:"id"`
	}

	EventUpdateResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		AuthorID    string `json:"author_id"`
		AuthorName  string `json:"author_name"`
		Price       int    `json:"price"`
		Capacity    int    `json:"capacity"`
		Availabilty int    `json:"availabilty"`
	}
)
