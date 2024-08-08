package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsPolicyStatement OBS桶存放的策略内容。
type ObsPolicyStatement struct {
	Effect *PolicyEffectEnum `json:"effect,omitempty"`

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
