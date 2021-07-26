package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAntitamperRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 单页条数

	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListAntitamperRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAntitamperRuleRequest struct{}"
	}

	return strings.Join([]string{"ListAntitamperRuleRequest", string(data)}, " ")
}
