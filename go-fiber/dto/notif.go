package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER           = "failed get list user"
	MESSAGE_FAILED_GET_LIST_EVENT          = "failed get list event"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get user token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get user"
	MESSAGE_FAILED_GET_EVENT               = "failed get event"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD = "wrong email or password"
	MESSAGE_FAILED_UPDATE_USER             = "failed update user"
	MESSAGE_FAILED_DELETE_USER             = "failed delete user"
	MESSAGE_FAILED_PROSES_REQUEST          = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS           = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL            = "failed verify email"
	MESSAGE_FAILED_CREATE_EVENT            = "failed to create event"
	MESSAGE_FAILED_GET_EVENT_BY_ID         = "failed to get event by this id"
	MESSAGE_FAILED_UPDATE_EVENT            = "failed to update event"
	MESSAGE_FAILED_DELETE_EVENT            = "failed to delete event"
	MESSAGE_FAILED_GET_AUTHOR              = "failed to get author"
	MESSAGE_FAILED_UNAUTHORIZED            = "failed to create event"
	MESSAGE_FAILED_GET_BUYER               = "failed to get buyer"
	MESSAGE_FAILED_CREATE_TRANSACTION      = "failed to create transaction"
	MESSAGE_FAILED_GET_LIST_TRANSACTION    = "failed to get list transaction"
	MESSAGE_FAILED_GET_TRANSACTION_BY_ID   = "failed to get transaction by id"
	MESSAGE_FAILED_UPDATE_TRANSACTION      = "failed to update transaction"
	MESSAGE_FAILED_DELETE_TRANSACTION      = "failed to delete transaction"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list user"
	MESSAGE_SUCCESS_GET_LIST_EVENT          = "success get list event"
	MESSAGE_SUCCESS_GET_USER                = "success get user"
	MESSAGE_SUCCESS_GET_EVENT               = "success get event"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update user"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete user"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
	MESSAGE_SUCCESS_CREATE_EVENT            = "success to create event"
	MESSAGE_SUCCESS_GET_EVENT_BY_ID         = "success to get event by this id"
	MESSAGE_SUCCESS_UPDATE_EVENT            = "success to update event"
	MESSAGE_SUCCESS_DELETE_EVENT            = "success to delete event"
	MESSAGE_SUCCESS_CREATE_TRANSACTION      = "success to create transaction"
	MESSAGE_SUCCESS_GET_LIST_TRANSACTION    = "success to create list transaction"
	MESSAGE_SUCCESS_GET_TRANSACTION_BY_ID   = "success to get transaction"
	MESSAGE_SUCCESS_DELETE_TRANSACTION      = "success to delete transaction"
	MESSAGE_SUCCESS_UPDATE_TRANSACTION      = "success to update transaction"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrGetAllUser             = errors.New("failed to get all user")
	ErrGetUserById            = errors.New("failed to get user by id")
	ErrGetUserByEmail         = errors.New("failed to get user by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUpdateUser             = errors.New("failed to update user")
	ErrUserNotAdmin           = errors.New("user not admin")
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteUser             = errors.New("failed to delete user")
	ErrPasswordNotMatch       = errors.New("password not match")
	ErrEmailOrPassword        = errors.New("wrong email or password")
	ErrAccountNotVerified     = errors.New("account not verified")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")

	// Event
	ErrCreateEvent         = errors.New("failed to create event")
	ErrGetEventById        = errors.New("failed to get event by id")
	ErrGetEventByAuthor    = errors.New("failed to get event by author")
	ErrEventNotFound       = errors.New("No Event Found")
	ErrUpdateEvent         = errors.New("failed to update event")
	ErrDeleteEvent         = errors.New("failed to delete event")
	ErrAuthorIDNotProvided = errors.New("author id is not provided")
	ErrInvalidEventID      = errors.New("invalid event id")
	ErrInvalidAuthorID     = errors.New("invalid author id")
	ErrUnauthorized        = errors.New("invalid role to create event")

	ErrInvalidTransactionID = errors.New("Invalid Transaction ID")
	ErrBuyerIDNotProvided   = errors.New("buyer ID not provided")

	ErrCreateTransaction   = errors.New("failed to create transaction")
	ErrGetTransactionById  = errors.New("failed to get transaction by id")
	ErrTransactionNotFound = errors.New("no transaction was found")
	ErrDeleteTransaction   = errors.New("failed to delete transaction")
	ErrBuyerNotFound       = errors.New("no buyer id was found")
)
