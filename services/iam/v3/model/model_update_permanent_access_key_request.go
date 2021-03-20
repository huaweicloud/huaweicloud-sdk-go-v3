package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePermanentAccessKeyRequest struct {
	AccessKey string `json:"access_key"`

	Body *UpdatePermanentAccessKeyRequestBody `json:"body,omitempty"`
}

func (o UpdatePermanentAccessKeyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePermanentAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePermanentAccessKeyRequest", string(data)}, " ")
}
