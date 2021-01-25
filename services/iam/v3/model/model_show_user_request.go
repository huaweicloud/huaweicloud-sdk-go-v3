package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowUserRequest struct {
	UserId string `json:"user_id"`
}

func (o ShowUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserRequest struct{}"
	}

	return strings.Join([]string{"ShowUserRequest", string(data)}, " ")
}
