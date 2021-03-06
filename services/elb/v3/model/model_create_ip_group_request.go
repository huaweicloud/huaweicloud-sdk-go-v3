package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateIpGroupRequest struct {
	Body *CreateIpGroupRequestBody `json:"body,omitempty"`
}

func (o CreateIpGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequest", string(data)}, " ")
}
