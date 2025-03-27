package dto

import (
	"mime/multipart"

	"github.com/tapeds/go-fiber-template/entity"
)

type (
	UserCreateRequest struct {
		Name       string                `json:"name" form:"name"`
		TelpNumber string                `json:"telp_number" form:"telp_number"`
		Email      string                `json:"email" form:"email"`
		Image      *multipart.FileHeader `json:"image" form:"image"`
		Password   string                `json:"password" form:"password"`
	}

	UserResponse struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		TelpNumber string `json:"telp_number"`
		Role       string `json:"role"`
		ImageUrl   string `json:"image_url"`
		IsVerified bool   `json:"is_verified"`
	}

	UserPaginationResponse struct {
		Data []UserResponse `json:"data"`
		PaginationResponse
	}

	GetAllUserRepositoryResponse struct {
		Users []entity.User
		PaginationResponse
	}

	UserUpdateRequest struct {
		Name       string `json:"name" form:"name"`
		TelpNumber string `json:"telp_number" form:"telp_number"`
		Email      string `json:"email" form:"email"`
	}

	UserUpdateResponse struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		TelpNumber string `json:"telp_number"`
		Role       string `json:"role"`
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	SendVerificationEmailRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
	}

	VerifyEmailRequest struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	VerifyEmailResponse struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}

	UpdateStatusIsVerifiedRequest struct {
		UserId     string `json:"user_id" form:"user_id" binding:"required"`
		IsVerified bool   `json:"is_verified" form:"is_verified"`
	}
)
