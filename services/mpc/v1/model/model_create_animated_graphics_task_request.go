package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAnimatedGraphicsTaskRequest struct {
	Body *CreateAnimatedGraphicsTaskReq `json:"body,omitempty"`
}

func (o CreateAnimatedGraphicsTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskRequest", string(data)}, " ")
}
