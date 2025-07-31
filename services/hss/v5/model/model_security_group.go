package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroup 安全组列表
type SecurityGroup struct {

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`

	// 安全组名称
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 安全组描述
	SecurityGroupDescription *string `json:"security_group_description,omitempty"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}
