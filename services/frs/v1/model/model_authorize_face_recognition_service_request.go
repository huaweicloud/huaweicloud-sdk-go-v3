package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AuthorizeFaceRecognitionServiceRequest struct {
}

func (o AuthorizeFaceRecognitionServiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AuthorizeFaceRecognitionServiceRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeFaceRecognitionServiceRequest", string(data)}, " ")
}
