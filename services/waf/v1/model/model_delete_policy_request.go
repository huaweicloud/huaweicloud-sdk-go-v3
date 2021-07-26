package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePolicyRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePolicyRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequest", string(data)}, " ")
}
