package models

type ResponseModel struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type ResponsePaginationModel struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    MetaModel   `json:"meta"`
}

type MetaModel struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

// error message
const (
	CreateEventError = "Create Event Fail"
	UpdateEventError = "Update Event Fail"
	DeleteEventError = "Delete Event Fail"
	CreateUserError  = "Create User Fail"

	InvalidRequestBodyError = "invalid request body"
	InvalidEventIDError     = "invalid event id"
)

// success message
const (
	CreateEventSuccess = "Create Event Success"
	UpdateEventSuccess = "Update Event Success"
	DeleteEventSuccess = "Delete Event Success"
	CreateUserSuccess  = "Create User Success"
)
