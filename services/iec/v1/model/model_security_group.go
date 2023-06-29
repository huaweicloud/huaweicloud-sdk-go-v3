package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroup 安全组数据对象
type SecurityGroup struct {

	// 安全组的ID。UUID
	Id *string `json:"id,omitempty"`

	// 安全组的名称。
	Name *string `json:"name,omitempty"`

	// 安全组的描述。
	Description *string `json:"description,omitempty"`

	// 安全组规则列表。
	SecurityGroupRules *[]SecurityGroupRule `json:"security_group_rules,omitempty"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}
