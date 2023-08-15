package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowSecurityGroupRequest Request Object
type NeutronShowSecurityGroupRequest struct {

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`
}

func (o NeutronShowSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupRequest", string(data)}, " ")
}
