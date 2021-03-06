package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateQueueRequest struct {
	Body *CreateQueueReq `json:"body,omitempty"`
}

func (o CreateQueueRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateQueueRequest struct{}"
	}

	return strings.Join([]string{"CreateQueueRequest", string(data)}, " ")
}
