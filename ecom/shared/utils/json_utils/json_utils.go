package json_utils

import (
	"encoding/json"
	"net/http"
)

type responseStruct struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

func JsonEncode(res http.ResponseWriter, v any) error {
	err := json.NewEncoder(res).Encode(v)
	if err != nil {
		return err
	}
	return nil
}

func JsonDecode(req *http.Request, v any) error {
	err := json.NewDecoder(req.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func JsonMarshaller(message string, status string, statusCode int) ([]byte, error) {
	errorResponse := responseStruct{message, status, statusCode}
	return json.Marshal(errorResponse)
}
