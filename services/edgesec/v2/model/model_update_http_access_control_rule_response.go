package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpAccessControlRuleResponse Response Object
type UpdateHttpAccessControlRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则所在策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则所在策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// 是否设定生效时间
	Time *bool `json:"time,omitempty"`

	// 生效时间
	Start *int64 `json:"start,omitempty"`

	// 失效时间
	Terminal *int64 `json:"terminal,omitempty"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 命中条件
	Conditions *[]HttpAccessControlRuleCondition `json:"conditions,omitempty"`

	Action *HttpRuleAction `json:"action,omitempty"`

	// 创建来源
	Producer       *int32 `json:"producer,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateHttpAccessControlRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpAccessControlRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpAccessControlRuleResponse", string(data)}, " ")
}
