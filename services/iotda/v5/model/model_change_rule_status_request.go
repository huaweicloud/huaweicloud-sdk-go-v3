package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeRuleStatusRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 规则Id

	RuleId string `json:"rule_id"`

	Body *RuleStatus `json:"body,omitempty"`
}

func (o ChangeRuleStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeRuleStatusRequest", string(data)}, " ")
}
