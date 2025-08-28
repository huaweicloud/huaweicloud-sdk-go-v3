package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventAttackTagResponse Response Object
type ShowEventAttackTagResponse struct {

	// **参数解释**: 攻击成功阶段的数量 **取值范围**: 最小值0，最大值2147483647
	AttackSuccessNum *int32 `json:"attack_success_num,omitempty"`

	// **参数解释**: 攻击尝试阶段的数量 **取值范围**: 最小值0，最大值2147483647
	AttackAttemptNum *int32 `json:"attack_attempt_num,omitempty"`

	// **参数解释**: 攻击被阻断阶段的数量 **取值范围**: 最小值0，最大值2147483647
	AttackBlockedNum *int32 `json:"attack_blocked_num,omitempty"`

	// **参数解释**: 异常行为阶段的数量 **取值范围**: 最小值0，最大值2147483647
	AbnormalBehaviorNum *int32 `json:"abnormal_behavior_num,omitempty"`

	// **参数解释**: 主机失陷阶段的数量 **取值范围**: 最小值0，最大值2147483647
	CollapsibleHostNum *int32 `json:"collapsible_host_num,omitempty"`

	// **参数解释**: 系统脆弱性阶段的数量 **取值范围**: 最小值0，最大值2147483647
	SystemVulnerabilityNum *int32 `json:"system_vulnerability_num,omitempty"`
	HttpStatusCode         int    `json:"-"`
}

func (o ShowEventAttackTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventAttackTagResponse struct{}"
	}

	return strings.Join([]string{"ShowEventAttackTagResponse", string(data)}, " ")
}
