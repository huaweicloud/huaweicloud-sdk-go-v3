package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePolicyRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`

	Body *UpdatePolicyRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}
