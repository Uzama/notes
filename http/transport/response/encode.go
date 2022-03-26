package response

import (
	"encoding/json"
)

type Encoder struct {
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Success interface{} `json:"success"`
}

func Encode(data interface{}, err interface{}, success interface{}) []byte {

	payload := Encoder{
		Data:    data,
		Error:   err,
		Success: success,
	}

	message, _ := json.Marshal(payload)

	return message
}
