package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListDomainPermissionsForGroupRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
}

func (o KeystoneListDomainPermissionsForGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListDomainPermissionsForGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListDomainPermissionsForGroupRequest", string(data)}, " ")
}
