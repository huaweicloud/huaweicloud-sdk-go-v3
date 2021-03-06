package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePoolRequest struct {
	Body *CreatePoolRequestBody `json:"body,omitempty"`
}

func (o CreatePoolRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePoolRequest struct{}"
	}

	return strings.Join([]string{"CreatePoolRequest", string(data)}, " ")
}
