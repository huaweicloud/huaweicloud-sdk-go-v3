package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRuleRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 规则ID，用于唯一标识一条规则，在创建规则时由物联网平台分配获得。

	RuleId string `json:"rule_id"`

	Body *Rule `json:"body,omitempty"`
}

func (o UpdateRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleRequest", string(data)}, " ")
}
