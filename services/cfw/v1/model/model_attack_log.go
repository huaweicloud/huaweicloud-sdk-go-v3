package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackLog struct {

	// **参数解释**： 动作 **取值范围**： 不涉及
	Action *string `json:"action,omitempty"`

	// **参数解释**： 应用 **取值范围**： 不涉及
	App *string `json:"app,omitempty"`

	// **参数解释**： 攻击规则 **取值范围**： 不涉及
	AttackRule *string `json:"attack_rule,omitempty"`

	// **参数解释**： 攻击规则ID **取值范围**： 不涉及
	AttackRuleId *string `json:"attack_rule_id,omitempty"`

	// **参数解释**： 攻击类型 **取值范围**： 不涉及
	AttackType *string `json:"attack_type,omitempty"`

	// **参数解释**： 攻击方向 **取值范围**： 不涉及
	Direction *string `json:"direction,omitempty"`

	// **参数解释**： 目的IP **取值范围**： 不涉及
	DstIp *string `json:"dst_ip,omitempty"`

	// **参数解释**： 目的端口 **取值范围**： 不涉及
	DstPort *int32 `json:"dst_port,omitempty"`

	// **参数解释**： 目的地区Id **取值范围**： 不涉及
	DstRegionId *string `json:"dst_region_id,omitempty"`

	// **参数解释**： 目的地区名称 **取值范围**： 不涉及
	DstRegionName *string `json:"dst_region_name,omitempty"`

	// **参数解释**： 目的省份ID **取值范围**： 不涉及
	DstProvinceId *string `json:"dst_province_id,omitempty"`

	// **参数解释**： 目的省份名称 **取值范围**： 不涉及
	DstProvinceName *string `json:"dst_province_name,omitempty"`

	// **参数解释**： 目的城市Id **取值范围**： 不涉及
	DstCityId *string `json:"dst_city_id,omitempty"`

	// **参数解释**： 目的城市名称 **取值范围**： 不涉及
	DstCityName *string `json:"dst_city_name,omitempty"`

	// **参数解释**： 发生时间 **取值范围**： 不涉及
	EventTime *int64 `json:"event_time,omitempty"`

	// **参数解释**： 危险等级 **取值范围**： 不涉及
	Level *string `json:"level,omitempty"`

	// **参数解释**： 协议 **取值范围**： 不涉及
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释**： 来源 **取值范围**： 不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释**： 源IP **取值范围**： 不涉及
	SrcIp *string `json:"src_ip,omitempty"`

	// **参数解释**： 真实IP **取值范围**： 不涉及
	RealIp *string `json:"real_ip,omitempty"`

	// **参数解释**： 标签 **取值范围**： 不涉及
	Tag *int32 `json:"tag,omitempty"`

	// **参数解释**： 源端口 **取值范围**： 不涉及
	SrcPort *int32 `json:"src_port,omitempty"`

	// **参数解释**： 源地区Id **取值范围**： 不涉及
	SrcRegionId *string `json:"src_region_id,omitempty"`

	// **参数解释**： 源地区名称 **取值范围**： 不涉及
	SrcRegionName *string `json:"src_region_name,omitempty"`

	// **参数解释**： 源省份Id **取值范围**： 不涉及
	SrcProvinceId *string `json:"src_province_id,omitempty"`

	// **参数解释**： 源省份名称 **取值范围**： 不涉及
	SrcProvinceName *string `json:"src_province_name,omitempty"`

	// **参数解释**： 源城市Id **取值范围**： 不涉及
	SrcCityId *string `json:"src_city_id,omitempty"`

	// **参数解释**： 源城市 **取值范围**： 不涉及
	SrcCityName *string `json:"src_city_name,omitempty"`

	// **参数解释**： VGW Id **取值范围**： 不涉及
	VgwId *string `json:"vgw_id,omitempty"`
}

func (o AttackLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackLog struct{}"
	}

	return strings.Join([]string{"AttackLog", string(data)}, " ")
}
