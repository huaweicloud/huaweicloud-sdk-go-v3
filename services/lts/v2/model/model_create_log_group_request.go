package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateLogGroupRequest struct {
	Body *CreateLogGroupParams `json:"body,omitempty"`
}

func (o CreateLogGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLogGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateLogGroupRequest", string(data)}, " ")
}
