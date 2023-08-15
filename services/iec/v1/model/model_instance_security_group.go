package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceSecurityGroup 边缘实例关联的安全组
type InstanceSecurityGroup struct {

	// 安全组名称或者UUID。
	Name *string `json:"name,omitempty"`

	// 实例使用安全组规则的ID。
	Id *string `json:"id,omitempty"`
}

func (o InstanceSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSecurityGroup struct{}"
	}

	return strings.Join([]string{"InstanceSecurityGroup", string(data)}, " ")
}
