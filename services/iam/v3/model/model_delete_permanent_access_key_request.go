package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePermanentAccessKeyRequest struct {
	AccessKey string `json:"access_key"`
}

func (o DeletePermanentAccessKeyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePermanentAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"DeletePermanentAccessKeyRequest", string(data)}, " ")
}
