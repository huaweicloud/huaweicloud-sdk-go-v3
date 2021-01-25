package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateBindingDeviceRequest struct {
	Body *BindMfaDevice `json:"body,omitempty"`
}

func (o CreateBindingDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateBindingDeviceRequest struct{}"
	}

	return strings.Join([]string{"CreateBindingDeviceRequest", string(data)}, " ")
}
