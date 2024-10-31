package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpRuleAction 防护规则动作
type HttpRuleAction struct {

	// 操作类型。   - “block”：拦截。   - “pass”：放行。   - “log”：仅记录
	Category string `json:"category"`

	// 攻击惩罚规则id，只有当category参数值为block时才可配置该参数
	FollowedActionId *string `json:"followed_action_id,omitempty"`

	Detail *HttpRuleActionDetail `json:"detail,omitempty"`
}

func (o HttpRuleAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpRuleAction struct{}"
	}

	return strings.Join([]string{"HttpRuleAction", string(data)}, " ")
}
