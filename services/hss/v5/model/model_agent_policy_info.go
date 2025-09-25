package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentPolicyInfo 策略信息
type AgentPolicyInfo struct {

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略组ID
	GroupId *string `json:"group_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// feature名称
	FeatureName *string `json:"feature_name,omitempty"`

	// 策略类别
	PolicyCategory *string `json:"policy_category,omitempty"`

	// **参数解释**: 策略应用状态 **约束限制**: 不涉及 **取值范围**: Agent的状态分为两类： - not_applied：未应用 - protection_degradation_not_applied：防护降级未应用 - processing：应用中 - applied：已应用 **默认取值**: 不涉及
	PolicyStatus *string `json:"policy_status,omitempty"`
}

func (o AgentPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentPolicyInfo struct{}"
	}

	return strings.Join([]string{"AgentPolicyInfo", string(data)}, " ")
}
