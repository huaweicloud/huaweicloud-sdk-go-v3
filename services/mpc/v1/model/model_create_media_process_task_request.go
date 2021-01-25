package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMediaProcessTaskRequest struct {
	Body *CreateMediaProcessReq `json:"body,omitempty"`
}

func (o CreateMediaProcessTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMediaProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessTaskRequest", string(data)}, " ")
}
