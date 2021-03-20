package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneAddUserToGroupRequest struct {
	GroupId string `json:"group_id"`

	UserId string `json:"user_id"`
}

func (o KeystoneAddUserToGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneAddUserToGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneAddUserToGroupRequest", string(data)}, " ")
}
