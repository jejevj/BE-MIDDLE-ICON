package repository

import (
	"context"
	"math"

	"github.com/google/uuid"
	"github.com/tapeds/go-fiber-template/dto"
	"github.com/tapeds/go-fiber-template/entity"
	"gorm.io/gorm"
)

type (
	TransactionRepository interface {
		CreateTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
		GetAllTransactionsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllTransactionRepositoryResponse, error)
		GetTransactionById(ctx context.Context, transactionId string) (entity.Transaction, error)
		UpdateTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
		DeleteTransaction(ctx context.Context, transactionId string) error
	}

	transactionRepository struct {
		db *gorm.DB
	}
)

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	tx := r.db

	// Ensure the buyer and event IDs are valid
	if transaction.EventID == "" {
		return entity.Transaction{}, dto.ErrBuyerIDNotProvided
	}

	if err := tx.WithContext(ctx).Create(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetAllTransactionsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllTransactionRepositoryResponse, error) {
	tx := r.db

	var transactions []entity.Transaction
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Transaction{}).Count(&count).Error; err != nil {
		return dto.GetAllTransactionRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&transactions).Error; err != nil {
		return dto.GetAllTransactionRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllTransactionRepositoryResponse{
		Transactions: transactions,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *transactionRepository) GetTransactionById(ctx context.Context, transactionId string) (entity.Transaction, error) {
	tx := r.db

	transactionUUID, err := uuid.Parse(transactionId)
	if err != nil {
		return entity.Transaction{}, dto.ErrInvalidTransactionID
	}

	var transaction entity.Transaction
	if err := tx.WithContext(ctx).Where("id = ?", transactionUUID).Take(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}

func (r *transactionRepository) UpdateTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	tx := r.db

	if transaction.EventID == "" {
		return entity.Transaction{}, dto.ErrBuyerIDNotProvided
	}

	if err := tx.WithContext(ctx).Updates(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}

func (r *transactionRepository) DeleteTransaction(ctx context.Context, transactionId string) error {
	tx := r.db

	transactionUUID, err := uuid.Parse(transactionId)
	if err != nil {
		return dto.ErrInvalidTransactionID
	}

	if err := tx.WithContext(ctx).Delete(&entity.Transaction{}, "id = ?", transactionUUID).Error; err != nil {
		return err
	}

	return nil
}
