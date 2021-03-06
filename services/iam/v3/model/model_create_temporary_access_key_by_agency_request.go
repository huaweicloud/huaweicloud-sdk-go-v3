package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTemporaryAccessKeyByAgencyRequest struct {
	Body *CreateTemporaryAccessKeyByAgencyRequestBody `json:"body,omitempty"`
}

func (o CreateTemporaryAccessKeyByAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByAgencyRequest", string(data)}, " ")
}
