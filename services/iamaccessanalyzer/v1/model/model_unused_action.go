package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnusedAction struct {

	// 授权项名称。
	Action string `json:"action"`

	// 用户使用授权项的最后访问时间。
	LastAccessed *sdktime.SdkTime `json:"last_accessed,omitempty"`

	// 通过该操作访问的角色来源列表。
	RoleSources *[]string `json:"role_sources,omitempty"`

	// 通过该操作访问的身份策略来源列表。
	IdentityPolicySources *[]string `json:"identity_policy_sources,omitempty"`
}

func (o UnusedAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedAction struct{}"
	}

	return strings.Join([]string{"UnusedAction", string(data)}, " ")
}
