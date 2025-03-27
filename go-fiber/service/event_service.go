package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tapeds/go-fiber-template/dto"
	"github.com/tapeds/go-fiber-template/entity"
	"github.com/tapeds/go-fiber-template/repository"
)

type (
	EventService interface {
		CreateEvent(ctx context.Context, req dto.EventCreateRequest) (dto.EventResponse, error)
		GetAllEventWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.EventPaginationResponse, error)
		GetEventById(ctx context.Context, eventId string) (dto.EventResponse, error)
		UpdateEvent(ctx context.Context, req dto.EventUpdateRequest, eventId string) (dto.EventUpdateResponse, error)
		DeleteEvent(ctx context.Context, eventId string) error
	}

	eventService struct {
		eventRepo  repository.EventRepository
		jwtService JWTService
	}
)

func NewEventService(eventRepo repository.EventRepository, jwtService JWTService) EventService {
	return &eventService{
		eventRepo:  eventRepo,
		jwtService: jwtService,
	}
}

func (s *eventService) CreateEvent(ctx context.Context, req dto.EventCreateRequest) (dto.EventResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	authorID, err := uuid.Parse(req.AuthorID)
	if err != nil {
		return dto.EventResponse{}, dto.ErrInvalidAuthorID
	}

	event := entity.Event{
		Name:        req.Name,
		AuthorID:    authorID,
		Price:       req.Price,
		Capacity:    req.Capacity,
		Availabilty: req.Availabilty,
	}

	eventReg, err := s.eventRepo.CreateEvent(ctx, event)
	if err != nil {
		return dto.EventResponse{}, dto.ErrCreateEvent
	}

	return dto.EventResponse{
		ID:          eventReg.ID.String(),
		Name:        eventReg.Name,
		AuthorID:    eventReg.AuthorID.String(),
		Price:       eventReg.Price,
		Capacity:    eventReg.Capacity,
		Availabilty: eventReg.Availabilty,
	}, nil
}

func (s *eventService) GetAllEventWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.EventPaginationResponse, error) {
	dataWithPaginate, err := s.eventRepo.GetAllEventWithPagination(ctx, req)
	if err != nil {
		return dto.EventPaginationResponse{}, err
	}

	var datas []dto.EventResponse
	for _, event := range dataWithPaginate.Events {
		data := dto.EventResponse{
			ID:          event.ID.String(),
			Name:        event.Name,
			AuthorID:    event.AuthorID.String(),
			Price:       event.Price,
			Capacity:    event.Capacity,
			Availabilty: event.Availabilty,
		}

		datas = append(datas, data)
	}

	return dto.EventPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *eventService) GetEventById(ctx context.Context, eventId string) (dto.EventResponse, error) {
	event, err := s.eventRepo.GetEventById(ctx, eventId)
	if err != nil {
		return dto.EventResponse{}, dto.ErrGetEventById
	}

	return dto.EventResponse{
		ID:          event.ID.String(),
		Name:        event.Name,
		AuthorID:    event.AuthorID.String(),
		Price:       event.Price,
		Capacity:    event.Capacity,
		Availabilty: event.Availabilty,
	}, nil
}

func (s *eventService) UpdateEvent(ctx context.Context, req dto.EventUpdateRequest, eventId string) (dto.EventUpdateResponse, error) {
	existingEvent, err := s.eventRepo.GetEventById(ctx, eventId)
	if err != nil {
		return dto.EventUpdateResponse{}, fmt.Errorf("failed to fetch event: %v", err)
	}

	capacityDifference := req.Capacity - existingEvent.Capacity

	newAvailability := existingEvent.Availabilty + capacityDifference

	updatedEvent := entity.Event{
		ID:          existingEvent.ID,
		Name:        req.Name,
		AuthorID:    existingEvent.AuthorID,
		Price:       req.Price,
		Capacity:    req.Capacity,
		Availabilty: newAvailability,
	}

	event, err := s.eventRepo.UpdateEvent(ctx, updatedEvent)
	if err != nil {
		return dto.EventUpdateResponse{}, fmt.Errorf("failed to update event: %v", err)
	}

	updatedEventDTO := dto.EventUpdateResponse{
		ID:          event.ID.String(),
		Name:        event.Name,
		AuthorID:    event.AuthorID.String(),
		Price:       event.Price,
		Capacity:    event.Capacity,
		Availabilty: event.Availabilty,
	}

	return updatedEventDTO, nil
}

func (s *eventService) DeleteEvent(ctx context.Context, eventId string) error {
	event, err := s.eventRepo.GetEventById(ctx, eventId)
	if err != nil {
		return dto.ErrEventNotFound
	}

	err = s.eventRepo.DeleteEvent(ctx, event.ID.String())
	if err != nil {
		return dto.ErrDeleteEvent
	}

	return nil
}
