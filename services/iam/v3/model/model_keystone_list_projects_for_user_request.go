package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListProjectsForUserRequest struct {
	UserId string `json:"user_id"`
}

func (o KeystoneListProjectsForUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListProjectsForUserRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListProjectsForUserRequest", string(data)}, " ")
}
