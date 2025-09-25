package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckRiskEventInfo 风险事件信息
type SecurityCheckRiskEventInfo struct {

	// **参数解释**： 风险级别 **取值范围**： - high：高危 - medium：中危 - low：低危 - safe：安全，无风险
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 告警事件名称 **取值范围**： 不涉及
	EventName *string `json:"event_name,omitempty"`

	// **参数解释**： 告警事件class，用于前台生成事件名称 **取值范围**： 不涉及
	EventClassId *string `json:"event_class_id,omitempty"`

	// **参数解释**： 攻击标识 **取值范围**： 不涉及
	AttackFlag *string `json:"attack_flag,omitempty"`

	// **参数解释**： 攻击时间 **取值范围**： 不涉及
	AttackTime *int64 `json:"attack_time,omitempty"`

	// **参数解释**： 处理状态 **取值范围**： - unhandled：未处理 - handled：已处理
	Status *string `json:"status,omitempty"`
}

func (o SecurityCheckRiskEventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckRiskEventInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckRiskEventInfo", string(data)}, " ")
}
