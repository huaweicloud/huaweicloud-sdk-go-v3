package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowServerGroupResponse struct {
	ServerGroup    *ShowServerGroupResult `json:"server_group,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowServerGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowServerGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowServerGroupResponse", string(data)}, " ")
}
