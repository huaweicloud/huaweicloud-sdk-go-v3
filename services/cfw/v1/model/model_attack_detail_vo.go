package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackDetailVo struct {

	// **参数解释**： 应用数量 **取值范围**： 不涉及
	AppCount *int64 `json:"app_count,omitempty"`

	// **参数解释**： 攻击规则数量 **取值范围**： 不涉及
	AttackRuleCount *int64 `json:"attack_rule_count,omitempty"`

	// **参数解释**： 攻击类型数量 **取值范围**： 不涉及
	AttackTypeCount *int64 `json:"attack_type_count,omitempty"`

	// **参数解释**： 攻击次数 **取值范围**： 不涉及
	Count *int64 `json:"count,omitempty"`

	// **参数解释**： 目的IP数量 **取值范围**： 不涉及
	DstIpCount *int64 `json:"dst_ip_count,omitempty"`

	// **参数解释**： 攻击端口数量 **取值范围**： 不涉及
	DstPortCount *int64 `json:"dst_port_count,omitempty"`

	// **参数解释**： 结束时间 **取值范围**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 攻击事件明细 **取值范围**： 不涉及
	Records *[]AttackLog `json:"records,omitempty"`

	// **参数解释**： 源IP数量 **取值范围**： 不涉及
	SrcIpCount *int64 `json:"src_ip_count,omitempty"`

	// **参数解释**： 开始时间 **取值范围**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int64 `json:"total,omitempty"`
}

func (o AttackDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackDetailVo struct{}"
	}

	return strings.Join([]string{"AttackDetailVo", string(data)}, " ")
}
