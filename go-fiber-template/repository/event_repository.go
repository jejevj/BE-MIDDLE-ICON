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
	EventRepository interface {
		CreateEvent(ctx context.Context, event entity.Event) (entity.Event, error)
		GetAllEventWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllEventRepositoryResponse, error)
		GetEventById(ctx context.Context, eventId string) (entity.Event, error)
		UpdateEvent(ctx context.Context, event entity.Event) (entity.Event, error)
		DeleteEvent(ctx context.Context, eventId string) error
	}

	eventRepository struct {
		db *gorm.DB
	}
)

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) CreateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	tx := r.db

	if event.AuthorID == uuid.Nil {
		return entity.Event{}, dto.ErrAuthorIDNotProvided
	}

	if err := tx.WithContext(ctx).Create(&event).Error; err != nil {
		return entity.Event{}, err
	}

	return event, nil
}

func (r *eventRepository) GetAllEventWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllEventRepositoryResponse, error) {
	tx := r.db

	var events []entity.Event
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Event{}).Count(&count).Error; err != nil {
		return dto.GetAllEventRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&events).Error; err != nil {
		return dto.GetAllEventRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllEventRepositoryResponse{
		Events: events,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *eventRepository) GetEventById(ctx context.Context, eventId string) (entity.Event, error) {
	tx := r.db

	eventUUID, err := uuid.Parse(eventId)
	if err != nil {
		return entity.Event{}, dto.ErrInvalidEventID
	}

	var event entity.Event
	if err := tx.WithContext(ctx).Where("id = ?", eventUUID).Take(&event).Error; err != nil {
		return entity.Event{}, err
	}

	return event, nil
}

func (r *eventRepository) UpdateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	tx := r.db

	if event.AuthorID == uuid.Nil {
		return entity.Event{}, dto.ErrAuthorIDNotProvided
	}

	if err := tx.WithContext(ctx).Updates(&event).Error; err != nil {
		return entity.Event{}, err
	}

	return event, nil
}

func (r *eventRepository) DeleteEvent(ctx context.Context, eventId string) error {
	tx := r.db

	eventUUID, err := uuid.Parse(eventId)
	if err != nil {
		return dto.ErrInvalidEventID
	}

	if err := tx.WithContext(ctx).Delete(&entity.Event{}, "id = ?", eventUUID).Error; err != nil {
		return err
	}

	return nil
}
