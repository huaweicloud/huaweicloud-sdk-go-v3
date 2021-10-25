package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateProtectionGroupRequest struct {
	Body *CreateProtectionGroupRequestBody `json:"body,omitempty"`
}

func (o CreateProtectionGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateProtectionGroupRequest", string(data)}, " ")
}
