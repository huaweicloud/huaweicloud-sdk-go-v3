package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSecurityGroupRequest struct {
	// 安全组的ID。

	SecurityGroupId string `json:"security_group_id"`
}

func (o ShowSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityGroupRequest", string(data)}, " ")
}
