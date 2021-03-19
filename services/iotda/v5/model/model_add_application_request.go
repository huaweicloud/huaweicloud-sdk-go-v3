package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddApplicationRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddApplication `json:"body,omitempty"`
}

func (o AddApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddApplicationRequest struct{}"
	}

	return strings.Join([]string{"AddApplicationRequest", string(data)}, " ")
}
