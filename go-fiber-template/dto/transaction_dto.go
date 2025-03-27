package dto

import (
	"github.com/tapeds/go-fiber-template/entity"
)

type (
	TransactionCreateRequest struct {
		EventID string `json:"event_id"`
		BuyerID string `json:"buyer_id"`
		Amount  int    `json:"amount"`
	}

	TransactionResponse struct {
		ID         string `json:"id"`
		BuyerID    string `json:"buyer_id"`
		BuyerName  string `json:"buyer_name"`
		BuyerEmail string `json:"buyer_email"`
		EventID    string `json:"event_id"`
		EventName  string `json:"event_name"`
		EventPrice int    `json:"event_price"`
		Amount     int    `json:"amount"`
		BuyedAt    string `json:"buyed_at"`
	}

	TransactionPaginationResponse struct {
		Data               []TransactionResponse `json:"data"`
		PaginationResponse `json:"meta"`
	}

	GetAllTransactionRepositoryResponse struct {
		Transactions []entity.Transaction
		PaginationResponse
	}

	TransactionUpdateRequest struct {
		Amount int `json:"amount"`
	}

	TransactionByIdRequest struct {
		ID string `json:"id"`
	}

	TransactionUpdateResponse struct {
		ID         string `json:"id"`
		BuyerName  string `json:"buyer_name"`
		BuyerEmail string `json:"buyer_email"`
		EventName  string `json:"event_name"`
		EventPrice int    `json:"event_price"`
		Amount     int    `json:"amount"`
		BuyedAt    string `json:"buyed_at"`
	}
)
