package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpAccessControlRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// 是否设定生效时间
	Time bool `json:"time"`

	// 生效时间
	Start *int64 `json:"start,omitempty"`

	// 失效时间
	Terminal *int64 `json:"terminal,omitempty"`

	// 优先级
	Priority int32 `json:"priority"`

	// 命中条件
	Conditions []HttpAccessControlRuleCondition `json:"conditions"`

	Action *HttpRuleAction `json:"action"`
}

func (o UpdateHttpAccessControlRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpAccessControlRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpAccessControlRuleRequestBody", string(data)}, " ")
}
