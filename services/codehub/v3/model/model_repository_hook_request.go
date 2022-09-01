package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryHookRequest struct {

	// 触发url
	HookUrl string `json:"hook_url" xml:"hook_url"`

	// 事件来源
	Service string `json:"service" xml:"service"`

	// 安全令牌
	Token *string `json:"token,omitempty" xml:"token"`

	// 触发事件
	HookEvents []string `json:"hook_events" xml:"hook_events"`
}

func (o RepositoryHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryHookRequest struct{}"
	}

	return strings.Join([]string{"RepositoryHookRequest", string(data)}, " ")
}
