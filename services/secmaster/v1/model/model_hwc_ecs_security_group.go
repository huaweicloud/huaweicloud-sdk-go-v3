package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcsSecurityGroup 安全组
type HwcEcsSecurityGroup struct {

	// 安全组ID。
	Id string `json:"id"`

	// 安全组名称或者UUID。
	Name *string `json:"name,omitempty"`
}

func (o HwcEcsSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcsSecurityGroup struct{}"
	}

	return strings.Join([]string{"HwcEcsSecurityGroup", string(data)}, " ")
}
