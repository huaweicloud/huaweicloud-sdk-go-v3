package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessLogInfo struct {

	// **参数解释**： 动作 **取值范围**： 不涉及
	Action *string `json:"action,omitempty"`

	// **参数解释**： 应用 **取值范围**： 不涉及
	App *string `json:"app,omitempty"`

	Url *string `json:"url,omitempty"`

	// **参数解释**： 目的域名 **取值范围**： 不涉及
	DstHost *string `json:"dst_host,omitempty"`

	// **参数解释**： 目的IP **取值范围**： 不涉及
	DstIp *string `json:"dst_ip,omitempty"`

	// **参数解释**： 目的端口 **取值范围**： 不涉及
	DstPort *int32 `json:"dst_port,omitempty"`

	// **参数解释**： 目的地区ID **取值范围**： 不涉及
	DstRegionId *string `json:"dst_region_id,omitempty"`

	// **参数解释**： 目的地区名称 **取值范围**： 不涉及
	DstRegionName *string `json:"dst_region_name,omitempty"`

	// **参数解释**： 目的省份ID **取值范围**： 不涉及
	DstProvinceId *string `json:"dst_province_id,omitempty"`

	// **参数解释**： 目的省份名称 **取值范围**： 不涉及
	DstProvinceName *string `json:"dst_province_name,omitempty"`

	// **参数解释**： 目的城市ID **取值范围**： 不涉及
	DstCityId *string `json:"dst_city_id,omitempty"`

	// **参数解释**： 目的城市名称 **取值范围**： 不涉及
	DstCityName *string `json:"dst_city_name,omitempty"`

	// **参数解释**： 命中时间 **取值范围**： 不涉及
	HitTime *int64 `json:"hit_time,omitempty"`

	LogId *string `json:"log_id,omitempty"`

	// **参数解释**： 协议 **取值范围**： 不涉及
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释**： 规则ID **取值范围**： 不涉及
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**： 规则名称 **取值范围**： 不涉及
	RuleName *string `json:"rule_name,omitempty"`

	RuleType *int32 `json:"rule_type,omitempty"`

	// **参数解释**： 源IP **取值范围**： 不涉及
	SrcIp *string `json:"src_ip,omitempty"`

	// **参数解释**： 源端口 **取值范围**： 不涉及
	SrcPort *int32 `json:"src_port,omitempty"`

	// **参数解释**： 源地区ID **取值范围**： 不涉及
	SrcRegionId *string `json:"src_region_id,omitempty"`

	// **参数解释**： 源地区名称 **取值范围**： 不涉及
	SrcRegionName *string `json:"src_region_name,omitempty"`

	// **参数解释**： 源省份ID **取值范围**： 不涉及
	SrcProvinceId *string `json:"src_province_id,omitempty"`

	// **参数解释**： 源省份名称 **取值范围**： 不涉及
	SrcProvinceName *string `json:"src_province_name,omitempty"`

	// **参数解释**： 源城市ID **取值范围**： 不涉及
	SrcCityId *string `json:"src_city_id,omitempty"`

	// **参数解释**： 源城市名称 **取值范围**： 不涉及
	SrcCityName *string `json:"src_city_name,omitempty"`

	// **参数解释**： VGW　ID **取值范围**： 不涉及
	VgwId *string `json:"vgw_id,omitempty"`

	QosRuleId *string `json:"qos_rule_id,omitempty"`

	QosRuleName *string `json:"qos_rule_name,omitempty"`

	QosRuleType *int32 `json:"qos_rule_type,omitempty"`
}

func (o AccessLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessLogInfo struct{}"
	}

	return strings.Join([]string{"AccessLogInfo", string(data)}, " ")
}
