package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneAddUserToGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneAddUserToGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneAddUserToGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneAddUserToGroupResponse", string(data)}, " ")
}
