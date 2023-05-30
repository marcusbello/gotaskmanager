package controllers

import (
	"github.com/marcusbello/gotaskmanager/models"
)

type (
	// For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	// For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	// Repsonse for autorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Model for authorized user with access token
	AuthUserModel struct {
		User  string `json:"user"`
		Token string `json:"token"`
	}
)
