package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsPolicyStatement obs桶存放的策略内容。
type ObsPolicyStatement struct {

	// 状态(正常、禁用)： * 'Allow' - 允许 * 'Deny' - 禁用
	Effect *string `json:"effect,omitempty"`

	// 可以进行操作的权限合集。
	Action *[]string `json:"action,omitempty"`

	// 允许访问的资源。
	Resource *[]string `json:"resource,omitempty"`
}

func (o ObsPolicyStatement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsPolicyStatement struct{}"
	}

	return strings.Join([]string{"ObsPolicyStatement", string(data)}, " ")
}
