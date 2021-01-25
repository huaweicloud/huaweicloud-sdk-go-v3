package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneRemoveDomainPermissionFromGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneRemoveDomainPermissionFromGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneRemoveDomainPermissionFromGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveDomainPermissionFromGroupResponse", string(data)}, " ")
}
