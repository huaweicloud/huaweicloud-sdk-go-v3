package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackTopRespBody struct {

	// **参数解释**： TOP攻击规则 **取值范围**： 不涉及
	AttackRule *[]TopInfo `json:"attack_rule,omitempty"`

	// **参数解释**： TOP攻击类型 **取值范围**： 不涉及
	AttackType *[]TopInfo `json:"attack_type,omitempty"`

	// **参数解释**： TOP攻击目的IP **取值范围**： 不涉及
	DstIp *[]TopInfo `json:"dst_ip,omitempty"`

	// **参数解释**： TOP被攻击端口 **取值范围**： 不涉及
	DstPort *[]TopInfo `json:"dst_port,omitempty"`

	// **参数解释**： TOP内部攻击来源IP **取值范围**： 不涉及
	InSrcIp *[]TopInfo `json:"in_src_ip,omitempty"`

	// **参数解释**： TOP威胁等级 **取值范围**： 不涉及
	Level *[]TopInfo `json:"level,omitempty"`

	// **参数解释**： TOP外部攻击来源IP **取值范围**： 不涉及
	OutSrcIp *[]TopInfo `json:"out_src_ip,omitempty"`

	// **参数解释**： TOP攻击来源IP **取值范围**： 不涉及
	SrcIp *[]TopInfo `json:"src_ip,omitempty"`

	// **参数解释**： 源地区ID **取值范围**： 不涉及
	SrcRegionId *[]TopInfo `json:"src_region_id,omitempty"`
}

func (o AttackTopRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackTopRespBody struct{}"
	}

	return strings.Join([]string{"AttackTopRespBody", string(data)}, " ")
}
