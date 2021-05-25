package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSecurityGroupRequest struct {
	// 安全组资源ID

	SecurityGroupId string `json:"security_group_id"`
}

func (o DeleteSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupRequest", string(data)}, " ")
}
