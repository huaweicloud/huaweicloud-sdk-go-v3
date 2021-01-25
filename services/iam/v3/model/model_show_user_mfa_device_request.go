package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowUserMfaDeviceRequest struct {
	UserId string `json:"user_id"`
}

func (o ShowUserMfaDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserMfaDeviceRequest struct{}"
	}

	return strings.Join([]string{"ShowUserMfaDeviceRequest", string(data)}, " ")
}
