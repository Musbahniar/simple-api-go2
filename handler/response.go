package handler

import (
	"encoding/json"
	"net/http"
	"simple-api-go2/logger"
)

//Response api object
type Response struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//Error api response Object
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHttpResponse(w http.ResponseWriter, status int, success bool, payload interface{}, httpErr interface{}) {
	response, err := json.Marshal(
		&Response{
			Success: success,
			Error:   httpErr,
			Data:    payload,
		},
	)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func HttpError(w http.ResponseWriter, status int, message string, details interface{}) {
	err := Error{
		Code:    status,
		Message: message,
	}
	logger.Errorf("Error While processing Request", err, details)
	NewHttpResponse(w, status, false, nil, err)
}

func HttpValidasi(w http.ResponseWriter, status int, message string, payload interface{}) {
	err := Error{
		Code:    status,
		Message: message,
	}
	NewHttpResponse(w, status, false, payload, err)
}

func HttpResponse(w http.ResponseWriter, status int, payload interface{}) {
	NewHttpResponse(w, status, true, payload, nil)
}

func HandleError(err error) {
	if err != nil {
		logger.Errorf("Error While processing Request", err, err.Error())
	}
}
