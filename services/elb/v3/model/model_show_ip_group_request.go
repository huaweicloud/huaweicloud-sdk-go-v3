package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowIpGroupRequest struct {
	IpgroupId string `json:"ipgroup_id"`
}

func (o ShowIpGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowIpGroupRequest", string(data)}, " ")
}
