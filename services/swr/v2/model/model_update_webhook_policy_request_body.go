package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWebhookPolicyRequestBody struct {

	// 触发器名称，由字母、汉字、数字、下划线（_）、中划线 (-)组成，1-256个字符。
	Name string `json:"name"`

	// 触发器描述
	Description *string `json:"description,omitempty"`

	// 触发器目标参数
	Targets []Target `json:"targets"`

	// 作用范围规则
	ScopeRules []ScopeRule `json:"scope_rules"`

	// 触发器触发条件，当前支持PUSH_ARTIFACT
	EventTypes *[]string `json:"event_types,omitempty"`

	// 是否使用
	Enabled bool `json:"enabled"`
}

func (o UpdateWebhookPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebhookPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWebhookPolicyRequestBody", string(data)}, " ")
}
