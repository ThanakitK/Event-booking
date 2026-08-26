package utils

import "event-booking/core/models"

func Response(code int, message string, data interface{}) models.ResponseModel {
	status := "success"
	if code >= 400 {
		status = "error"
	}
	return models.ResponseModel{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ResponsePagination(code int, message string, data interface{}, meta models.MetaModel) models.ResponsePaginationModel {
	status := "success"
	if code >= 400 {
		status = "error"
	}
	return models.ResponsePaginationModel{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
