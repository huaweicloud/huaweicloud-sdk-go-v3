package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneRemoveProjectPermissionFromGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneRemoveProjectPermissionFromGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneRemoveProjectPermissionFromGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveProjectPermissionFromGroupResponse", string(data)}, " ")
}
