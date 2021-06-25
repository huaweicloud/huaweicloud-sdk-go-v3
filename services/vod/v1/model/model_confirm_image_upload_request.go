package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ConfirmImageUploadRequest struct {
	Body *ConfirmImageUploadReq `json:"body,omitempty"`
}

func (o ConfirmImageUploadRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfirmImageUploadRequest struct{}"
	}

	return strings.Join([]string{"ConfirmImageUploadRequest", string(data)}, " ")
}
