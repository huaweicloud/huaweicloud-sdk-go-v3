package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackReport struct {

	// **参数解释**： TOP攻击目的IP **取值范围**： 不涉及
	DstIp *[]ItemVo `json:"dst_ip,omitempty"`

	// **参数解释**： 入侵防御状态 **取值范围**： 不涉及
	IpsMode *int32 `json:"ips_mode,omitempty"`

	// **参数解释**： 攻击等级分布 **取值范围**： 不涉及
	Level *[]ItemVo `json:"level,omitempty"`

	// **参数解释**： TOP攻击规则 **取值范围**： 不涉及
	Rule *[]ItemVo `json:"rule,omitempty"`

	// **参数解释**： TOP来源IP **取值范围**： 不涉及
	SrcIp *[]ItemVo `json:"src_ip,omitempty"`

	// **参数解释**： 攻击趋势 **取值范围**： 不涉及
	Trend *[]TrendVo `json:"trend,omitempty"`

	// **参数解释**： TOP攻击分布 **取值范围**： 不涉及
	Type *[]ItemVo `json:"type,omitempty"`
}

func (o AttackReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackReport struct{}"
	}

	return strings.Join([]string{"AttackReport", string(data)}, " ")
}
