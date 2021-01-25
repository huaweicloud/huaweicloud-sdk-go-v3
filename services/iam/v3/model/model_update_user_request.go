package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateUserRequest struct {
	UserId string                 `json:"user_id"`
	Body   *UpdateUserRequestBody `json:"body,omitempty"`
}

func (o UpdateUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserRequest", string(data)}, " ")
}
