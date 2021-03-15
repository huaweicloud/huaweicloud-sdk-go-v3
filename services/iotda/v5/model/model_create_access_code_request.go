package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAccessCodeRequest struct {
	InstanceId *string                      `json:"Instance-Id,omitempty"`
	Body       *CreateAccessCodeRequestBody `json:"body,omitempty"`
}

func (o CreateAccessCodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAccessCodeRequest struct{}"
	}

	return strings.Join([]string{"CreateAccessCodeRequest", string(data)}, " ")
}
