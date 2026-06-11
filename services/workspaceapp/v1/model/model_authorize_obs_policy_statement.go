package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeObsPolicyStatement OBS桶存放的策略内容。
type AuthorizeObsPolicyStatement struct {
	Effect *PolicyEffectEnum `json:"effect,omitempty"`

	// 可以进行操作的权限合集。
	Action *[]string `json:"action,omitempty"`

	// 允许访问的资源。
	Resource *[]string `json:"resource,omitempty"`
}

func (o AuthorizeObsPolicyStatement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeObsPolicyStatement struct{}"
	}

	return strings.Join([]string{"AuthorizeObsPolicyStatement", string(data)}, " ")
}
