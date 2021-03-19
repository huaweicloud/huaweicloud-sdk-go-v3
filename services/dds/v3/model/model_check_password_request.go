package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckPasswordRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CheckPasswordRequestBody `json:"body,omitempty"`
}

func (o CheckPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckPasswordRequest struct{}"
	}

	return strings.Join([]string{"CheckPasswordRequest", string(data)}, " ")
}
