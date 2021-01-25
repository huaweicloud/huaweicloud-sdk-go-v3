package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCheckDomainPermissionForGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneCheckDomainPermissionForGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckDomainPermissionForGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCheckDomainPermissionForGroupResponse", string(data)}, " ")
}
