package service

import (
	"context"
	"fmt"

	"github.com/tapeds/go-fiber-template/dto"
	"github.com/tapeds/go-fiber-template/entity"
	"github.com/tapeds/go-fiber-template/repository"
)

type (
	TransactionService interface {
		CreateTransaction(ctx context.Context, req dto.TransactionCreateRequest) (dto.TransactionResponse, error)
		GetAllTransactionsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TransactionPaginationResponse, error)
		GetTransactionById(ctx context.Context, transactionId string) (dto.TransactionResponse, error)
		UpdateTransaction(ctx context.Context, req dto.TransactionUpdateRequest, transactionId string) (dto.TransactionUpdateResponse, error)
		DeleteTransaction(ctx context.Context, transactionId string) error
	}

	transactionService struct {
		transactionRepo repository.TransactionRepository
		eventRepo       repository.EventRepository
		userRepo        repository.UserRepository
	}
)

func NewTransactionService(transactionRepo repository.TransactionRepository, eventRepo repository.EventRepository, userRepo repository.UserRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		eventRepo:       eventRepo,
		userRepo:        userRepo,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req dto.TransactionCreateRequest) (dto.TransactionResponse, error) {
	// Ensure valid BuyerID and EventID
	event, err := s.eventRepo.GetEventById(ctx, req.EventID)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrEventNotFound
	}

	buyer, err := s.userRepo.GetUserById(ctx, req.BuyerID)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrBuyerNotFound
	}

	transaction := entity.Transaction{
		EventID: req.EventID,
		BuyerID: req.BuyerID,
		Amount:  req.Amount,
	}

	transactionReg, err := s.transactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrCreateTransaction
	}

	return dto.TransactionResponse{
		ID:         transactionReg.ID.String(),
		BuyerName:  buyer.Name,
		BuyerEmail: buyer.Email,
		EventName:  event.Name,
		EventPrice: event.Price,
		Amount:     transactionReg.Amount,
		BuyedAt:    transactionReg.CreatedAt.String(),
	}, nil
}

func (s *transactionService) GetAllTransactionsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TransactionPaginationResponse, error) {
	dataWithPaginate, err := s.transactionRepo.GetAllTransactionsWithPagination(ctx, req)
	if err != nil {
		return dto.TransactionPaginationResponse{}, err
	}

	var datas []dto.TransactionResponse
	for _, transaction := range dataWithPaginate.Transactions {
		event, err := s.eventRepo.GetEventById(ctx, transaction.EventID)
		if err != nil {
			return dto.TransactionPaginationResponse{}, err
		}

		buyer, err := s.userRepo.GetUserById(ctx, transaction.BuyerID)
		if err != nil {
			return dto.TransactionPaginationResponse{}, err
		}

		data := dto.TransactionResponse{
			ID:         transaction.ID.String(),
			BuyerID:    buyer.ID.String(),
			BuyerName:  buyer.Name,
			BuyerEmail: buyer.Email,
			EventID:    event.ID.String(),
			EventName:  event.Name,
			EventPrice: event.Price,
			Amount:     transaction.Amount,
			BuyedAt:    transaction.CreatedAt.String(),
		}

		datas = append(datas, data)
	}

	return dto.TransactionPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *transactionService) GetTransactionById(ctx context.Context, transactionId string) (dto.TransactionResponse, error) {
	transaction, err := s.transactionRepo.GetTransactionById(ctx, transactionId)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrGetTransactionById
	}

	event, err := s.eventRepo.GetEventById(ctx, transaction.EventID)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	buyer, err := s.userRepo.GetUserById(ctx, transaction.BuyerID)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	return dto.TransactionResponse{
		ID:         transaction.ID.String(),
		BuyerID:    buyer.ID.String(),
		BuyerName:  buyer.Name,
		BuyerEmail: buyer.Email,
		EventID:    event.ID.String(),
		EventName:  event.Name,
		EventPrice: event.Price,
		Amount:     transaction.Amount,
		BuyedAt:    transaction.CreatedAt.String(),
	}, nil
}

func (s *transactionService) UpdateTransaction(ctx context.Context, req dto.TransactionUpdateRequest, transactionId string) (dto.TransactionUpdateResponse, error) {
	transaction, err := s.transactionRepo.GetTransactionById(ctx, transactionId)
	if err != nil {
		return dto.TransactionUpdateResponse{}, fmt.Errorf("failed to fetch transaction: %v", err)
	}

	transaction.Amount = req.Amount

	updatedTransaction, err := s.transactionRepo.UpdateTransaction(ctx, transaction)
	if err != nil {
		return dto.TransactionUpdateResponse{}, fmt.Errorf("failed to update transaction: %v", err)
	}

	event, err := s.eventRepo.GetEventById(ctx, updatedTransaction.EventID)
	if err != nil {
		return dto.TransactionUpdateResponse{}, err
	}

	buyer, err := s.userRepo.GetUserById(ctx, updatedTransaction.BuyerID)
	if err != nil {
		return dto.TransactionUpdateResponse{}, err
	}

	return dto.TransactionUpdateResponse{
		ID:         updatedTransaction.ID.String(),
		BuyerName:  buyer.Name,
		BuyerEmail: buyer.Email,
		EventName:  event.Name,
		EventPrice: event.Price,
		Amount:     updatedTransaction.Amount,
		BuyedAt:    updatedTransaction.CreatedAt.String(),
	}, nil
}

func (s *transactionService) DeleteTransaction(ctx context.Context, transactionId string) error {
	transaction, err := s.transactionRepo.GetTransactionById(ctx, transactionId)
	if err != nil {
		return dto.ErrTransactionNotFound
	}

	err = s.transactionRepo.DeleteTransaction(ctx, transaction.ID.String())
	if err != nil {
		return dto.ErrDeleteTransaction
	}

	return nil
}
