package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCustomRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 匹配条件列表，匹配条件必须同时满足。
	Conditions *[]CustomRuleConditions `json:"conditions,omitempty"`

	Action *CustomAction `json:"action,omitempty"`

	// 预留参数，可忽略。
	ActionMode *bool `json:"action_mode,omitempty"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：0到1000。
	Priority *int32 `json:"priority,omitempty"`

	// 创建精准防护规则的
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 精准防护规则生效时间:  - “false”：表示该规则立即生效。   - “true”：表示自定义生效时间。
	Time *bool `json:"time,omitempty"`

	// 精准防护规则生效的起始时间戳（毫秒）。当time=true，才会返回该参数。
	Start *int64 `json:"start,omitempty"`

	// 精准防护规则生效的终止时间戳（毫秒）。当time=true，才会返回该参数。
	Terminal *int64 `json:"terminal,omitempty"`

	// 规则创建对象，该参数为预留参数，用于后续功能扩展，当前请用户忽略该参数
	Producer       *int32 `json:"producer,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateCustomRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomRuleResponse", string(data)}, " ")
}
