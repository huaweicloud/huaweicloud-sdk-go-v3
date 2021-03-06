package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTemporaryAccessKeyByTokenRequest struct {
	Body *CreateTemporaryAccessKeyByTokenRequestBody `json:"body,omitempty"`
}

func (o CreateTemporaryAccessKeyByTokenRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByTokenRequest", string(data)}, " ")
}
