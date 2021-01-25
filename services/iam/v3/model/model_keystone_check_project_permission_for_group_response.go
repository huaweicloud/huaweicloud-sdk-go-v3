package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCheckProjectPermissionForGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneCheckProjectPermissionForGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckProjectPermissionForGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCheckProjectPermissionForGroupResponse", string(data)}, " ")
}
