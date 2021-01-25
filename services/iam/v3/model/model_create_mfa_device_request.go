package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMfaDeviceRequest struct {
	Body *CreateMfaDeviceReq `json:"body,omitempty"`
}

func (o CreateMfaDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMfaDeviceRequest struct{}"
	}

	return strings.Join([]string{"CreateMfaDeviceRequest", string(data)}, " ")
}
