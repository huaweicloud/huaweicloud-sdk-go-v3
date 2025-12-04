package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomRuleResponse Response Object
type UpdateCustomRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 匹配条件列表，匹配条件必须同时满足。
	Conditions *[]CustomRuleConditions `json:"conditions,omitempty"`

	Action *CustomAction `json:"action,omitempty"`

	// 预留参数，可忽略。
	ActionMode *bool `json:"action_mode,omitempty"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：0到65535。
	Priority *int32 `json:"priority,omitempty"`

	// 精准防护规则生效时间:  - “false”：表示该规则立即生效。   - “true”：表示自定义生效时间。
	Time *bool `json:"time,omitempty"`

	// 精准防护规则生效的起始时间戳（秒）。当time=true，才会返回该参数。
	Start *int64 `json:"start,omitempty"`

	// 精准防护规则生效的终止时间戳（秒）。当time=true，才会返回该参数。
	Terminal *int64 `json:"terminal,omitempty"`

	// 规则创建对象，该参数为预留参数，用于后续功能扩展，当前请用户忽略该参数
	Producer       *int32 `json:"producer,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateCustomRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateCustomRuleResponse", string(data)}, " ")
}
