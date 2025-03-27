package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tapeds/go-fiber-template/constants"
	"github.com/tapeds/go-fiber-template/dto"
	"github.com/tapeds/go-fiber-template/service"
	"github.com/tapeds/go-fiber-template/utils"
)

type (
	EventController interface {
		CreateEvent(ctx *fiber.Ctx) error
		GetAllEvent(ctx *fiber.Ctx) error
		GetEventById(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	eventController struct {
		eventService service.EventService
		userService  service.UserService
	}
)

func NewEventController(eventService service.EventService, userService service.UserService) EventController {
	return &eventController{
		eventService: eventService,
		userService:  userService,
	}
}

func (c *eventController) CreateEvent(ctx *fiber.Ctx) error {
	var req dto.EventCreateRequest

	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	userId := ctx.Locals("user_id").(string)

	author, err := c.userService.GetUserById(ctx.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_AUTHOR, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if author.Role != constants.ENUM_ROLE_ADMIN {
		res := utils.BuildResponseFailed("Unauthorized", dto.MESSAGE_FAILED_UNAUTHORIZED, nil)
		return ctx.Status(http.StatusForbidden).JSON(res)
	}

	req.AuthorID = userId

	result, err := c.eventService.CreateEvent(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_EVENT, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_EVENT, struct {
		*dto.EventResponse
		AuthorName string `json:"author_name"`
	}{
		EventResponse: &dto.EventResponse{
			ID:          result.ID,
			Name:        result.Name,
			AuthorID:    result.AuthorID,
			Price:       result.Price,
			Capacity:    result.Capacity,
			Availabilty: result.Availabilty,
		},
		AuthorName: author.Name,
	})

	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *eventController) GetAllEvent(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	// Parse request body into PaginationRequest DTO
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Get paginated list of events from service
	result, err := c.eventService.GetAllEventWithPagination(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_EVENT, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	var eventsWithAuthorName []dto.EventResponse
	for _, event := range result.Data {
		author, err := c.userService.GetUserById(ctx.Context(), event.AuthorID)
		if err != nil {
			author.Name = "Unknown"
		}

		event.AuthorName = author.Name
		eventsWithAuthorName = append(eventsWithAuthorName, event)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_EVENT,
		Data:    eventsWithAuthorName,
		Meta:    result.PaginationResponse,
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *eventController) GetEventById(ctx *fiber.Ctx) error {
	var req dto.EventByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.eventService.GetEventById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_EVENT_BY_ID, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	author, err := c.userService.GetUserById(ctx.Context(), result.AuthorID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_AUTHOR, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_EVENT_BY_ID,
		Data: struct {
			*dto.EventResponse
			AuthorName string `json:"author_name"`
		}{
			EventResponse: &dto.EventResponse{
				ID:          result.ID,
				Name:        result.Name,
				AuthorID:    result.AuthorID,
				Price:       result.Price,
				Capacity:    result.Capacity,
				Availabilty: result.Availabilty,
			},
			AuthorName: author.Name,
		},
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *eventController) Update(ctx *fiber.Ctx) error {
	var req dto.EventUpdateRequest

	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("Event ID is required", "Event ID is missing or invalid.", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.eventService.UpdateEvent(ctx.Context(), req, req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_EVENT, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_EVENT, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *eventController) Delete(ctx *fiber.Ctx) error {
	var req dto.EventByIdRequest

	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("Event ID is required", "Event ID is missing or invalid.", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if err := c.eventService.DeleteEvent(ctx.Context(), req.ID); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_EVENT, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_EVENT, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
