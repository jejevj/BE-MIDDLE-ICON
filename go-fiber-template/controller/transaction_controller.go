package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tapeds/go-fiber-template/dto"
	"github.com/tapeds/go-fiber-template/service"
	"github.com/tapeds/go-fiber-template/utils"
)

type (
	TransactionController interface {
		CreateTransaction(ctx *fiber.Ctx) error
		GetAllTransactions(ctx *fiber.Ctx) error
		GetTransactionById(ctx *fiber.Ctx) error
		UpdateTransaction(ctx *fiber.Ctx) error
		DeleteTransaction(ctx *fiber.Ctx) error
	}

	transactionController struct {
		transactionService service.TransactionService
		userService        service.UserService
		eventService       service.EventService
	}
)

func NewTransactionController(transactionService service.TransactionService, userService service.UserService, eventService service.EventService) TransactionController {
	return &transactionController{
		transactionService: transactionService,
		userService:        userService,
		eventService:       eventService,
	}
}

func (c *transactionController) CreateTransaction(ctx *fiber.Ctx) error {
	var req dto.TransactionCreateRequest

	// Parse the request body
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	userId := ctx.Locals("user_id").(string)

	// Validate Buyer
	_, err := c.userService.GetUserById(ctx.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BUYER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Validate Event
	event, err := c.eventService.GetEventById(ctx.Context(), req.EventID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_EVENT, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Check if event has enough availability
	if event.Availabilty < req.Amount {
		res := utils.BuildResponseFailed("Insufficient availability", "The event doesn't have enough spots available.", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Decrease the event availability by the amount of the transaction
	event.Availabilty -= req.Amount

	// Prepare the EventUpdateRequest
	eventUpdateRequest := dto.EventUpdateRequest{
		ID:          event.ID,          // Keep the event name as it is
		Name:        event.Name,        // Keep the event name as it is
		Price:       event.Price,       // Keep the price as is
		Capacity:    event.Capacity,    // Keep the capacity as it is
		Availabilty: event.Availabilty, // Update the availability
	}

	// Update the event's availability
	_, err = c.eventService.UpdateEvent(ctx.Context(), eventUpdateRequest, event.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_EVENT, err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	// Set the buyer ID for the transaction
	req.BuyerID = userId

	// Create the transaction
	result, err := c.transactionService.CreateTransaction(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TRANSACTION, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Fetch Event and Buyer details for the response
	event, _ = c.eventService.GetEventById(ctx.Context(), req.EventID)
	buyer, _ := c.userService.GetUserById(ctx.Context(), userId)

	// Send response with transaction details
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TRANSACTION, struct {
		*dto.TransactionResponse
		EventName  string `json:"event_name"`
		BuyerName  string `json:"buyer_name"`
		BuyerEmail string `json:"buyer_email"`
	}{
		TransactionResponse: &dto.TransactionResponse{
			ID:         result.ID,
			BuyerID:    buyer.ID,
			BuyerName:  buyer.Name,
			BuyerEmail: buyer.Email,
			EventID:    result.EventID,
			EventName:  event.Name,
			EventPrice: event.Price,
			Amount:     result.Amount,
			BuyedAt:    result.BuyedAt,
		},
		EventName:  event.Name,
		BuyerName:  buyer.Name,
		BuyerEmail: buyer.Email,
	})

	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *transactionController) GetAllTransactions(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.transactionService.GetAllTransactionsWithPagination(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TRANSACTION, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	var transactionsWithDetails []dto.TransactionResponse
	for _, transaction := range result.Data {
		// Fetch event and buyer details for each transaction
		event, err := c.eventService.GetEventById(ctx.Context(), transaction.EventID)
		if err != nil {
			event.Name = "Unknown"
		} else {
			transaction.EventName = event.Name
		}

		buyer, err := c.userService.GetUserById(ctx.Context(), transaction.BuyerID)
		if err != nil {
			buyer.Name = "Unknown"
		}

		// transaction.EventID = event.ID
		transaction.EventID = transaction.EventID
		transaction.BuyerName = buyer.Name
		transactionsWithDetails = append(transactionsWithDetails, transaction)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_TRANSACTION,
		Data:    transactionsWithDetails,
		Meta:    result.PaginationResponse,
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *transactionController) GetTransactionById(ctx *fiber.Ctx) error {
	var req dto.TransactionByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.transactionService.GetTransactionById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TRANSACTION_BY_ID, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	event, err := c.eventService.GetEventById(ctx.Context(), result.EventID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_EVENT, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	buyer, err := c.userService.GetUserById(ctx.Context(), result.BuyerID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BUYER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_TRANSACTION_BY_ID,
		Data: struct {
			*dto.TransactionResponse
			EventName  string `json:"event_name"`
			BuyerName  string `json:"buyer_name"`
			BuyerEmail string `json:"buyer_email"`
		}{
			TransactionResponse: &dto.TransactionResponse{
				ID:         result.ID,
				BuyerID:    result.BuyerID,
				EventID:    result.EventID,
				BuyerName:  buyer.Name,
				BuyerEmail: buyer.Email,
				EventName:  event.Name,
				EventPrice: event.Price,
				Amount:     result.Amount,
				BuyedAt:    result.BuyedAt,
			},
			EventName:  event.Name,
			BuyerName:  buyer.Name,
			BuyerEmail: buyer.Email,
		},
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *transactionController) UpdateTransaction(ctx *fiber.Ctx) error {
	var req dto.TransactionUpdateRequest

	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.Amount <= 0 {
		res := utils.BuildResponseFailed("Invalid Amount", "Transaction amount must be greater than zero.", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	transactionId := ctx.Params("id")
	result, err := c.transactionService.UpdateTransaction(ctx.Context(), req, transactionId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TRANSACTION, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TRANSACTION, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *transactionController) DeleteTransaction(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("id")

	if err := c.transactionService.DeleteTransaction(ctx.Context(), transactionId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TRANSACTION, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_TRANSACTION, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
