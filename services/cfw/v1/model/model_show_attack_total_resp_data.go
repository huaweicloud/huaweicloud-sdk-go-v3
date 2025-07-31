package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttackTotalRespData **参数解释**： 攻击统计信息 **取值范围**： 不涉及
type ShowAttackTotalRespData struct {

	// **参数解释**： 攻击次数 **取值范围**： 不涉及
	AttackCount *int64 `json:"attack_count,omitempty"`

	// **参数解释**： 阻断次数 **取值范围**： 不涉及
	DenyCount *int64 `json:"deny_count,omitempty"`

	// **参数解释**： 放行次数 **取值范围**： 不涉及
	PermitCount *int64 `json:"permit_count,omitempty"`

	// **参数解释**： 风险端口 **取值范围**： 不涉及
	RiskPorts *int64 `json:"risk_ports,omitempty"`
}

func (o ShowAttackTotalRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackTotalRespData struct{}"
	}

	return strings.Join([]string{"ShowAttackTotalRespData", string(data)}, " ")
}
