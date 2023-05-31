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
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
	// For Post/Put - /tasks
	// For Get -/tasks/{id}
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	// For Post/Put - /tasks
	// For Get -/tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
	// For Post/Put - /notes
	NoteResource struct {
		Data NoteModel `json:"data"`
	}
	// For Get - /notes
	// For /notes/tasks/id
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}
	//Model for a TaskNote
	NoteModel struct {
		TaskId      string `json:"taskid"`
		Description string `json:"description"`
	}
)
