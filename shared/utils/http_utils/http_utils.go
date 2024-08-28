package http_utils

import (
	"github.com/karthik-code78/ecom/shared/utils/json_utils"
	"net/http"
)

func SendErrorResponse(res http.ResponseWriter, message string, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	errorResponse, err := json_utils.JsonMarshaller(message, "error", statusCode)
	if err != nil {
		http.Error(res, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	_, err = res.Write(errorResponse)
	if err != nil {
		http.Error(res, "Error writing JSON response", http.StatusInternalServerError)
	}
}

func SendSuccessResponse(res http.ResponseWriter, message string, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	successResponse, err := json_utils.JsonMarshaller(message, "success", statusCode)
	if err != nil {
		http.Error(res, "Error generating JSON response", http.StatusInternalServerError)
	}
	_, err = res.Write(successResponse)
	if err != nil {
		http.Error(res, "Error writing JSON response", http.StatusInternalServerError)
	}
}
