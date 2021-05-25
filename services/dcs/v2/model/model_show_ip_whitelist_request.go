package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowIpWhitelistRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowIpWhitelistRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowIpWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowIpWhitelistRequest", string(data)}, " ")
}
