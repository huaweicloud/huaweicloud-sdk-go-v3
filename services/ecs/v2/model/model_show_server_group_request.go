package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowServerGroupRequest struct {
	ServerGroupId string `json:"server_group_id"`
}

func (o ShowServerGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowServerGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowServerGroupRequest", string(data)}, " ")
}
