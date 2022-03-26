package errors

import (
	"net/http"
	"notes/http/transport/response"
)

func HandleError(w http.ResponseWriter, err error, code int) {

	payload := response.Encode(nil, err.Error(), "false")

	response.Send(w, payload, code)
}
