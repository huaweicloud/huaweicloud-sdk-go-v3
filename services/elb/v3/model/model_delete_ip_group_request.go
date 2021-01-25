package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteIpGroupRequest struct {
	IpgroupId string `json:"ipgroup_id"`
}

func (o DeleteIpGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteIpGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupRequest", string(data)}, " ")
}
