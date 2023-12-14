package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionConstraint 版本限制说明
type VersionConstraint struct {

	// 其他限制
	Other map[string]interface{} `json:"other,omitempty"`

	NodeConstraint *NodeConstraints `json:"node_constraint,omitempty"`

	// 安全模式kerberos排除组件列表
	SafeModeKerberosExcludeComponents *[]string `json:"safe_mode_kerberos_exclude_components,omitempty"`
}

func (o VersionConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionConstraint struct{}"
	}

	return strings.Join([]string{"VersionConstraint", string(data)}, " ")
}
