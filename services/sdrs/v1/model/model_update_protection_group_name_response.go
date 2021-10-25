package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateProtectionGroupNameResponse struct {
	ServerGroup    *ShowProtectionGroupParams `json:"server_group,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o UpdateProtectionGroupNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProtectionGroupNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateProtectionGroupNameResponse", string(data)}, " ")
}
