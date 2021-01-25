package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListProjectPermissionsForGroupRequest struct {
	ProjectId string `json:"project_id"`
	GroupId   string `json:"group_id"`
}

func (o KeystoneListProjectPermissionsForGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListProjectPermissionsForGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListProjectPermissionsForGroupRequest", string(data)}, " ")
}
