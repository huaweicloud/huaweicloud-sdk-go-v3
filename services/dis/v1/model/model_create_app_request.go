package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAppRequest struct {
	Body *CreateAppRequestBody `json:"body,omitempty"`
}

func (o CreateAppRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAppRequest struct{}"
	}

	return strings.Join([]string{"CreateAppRequest", string(data)}, " ")
}
