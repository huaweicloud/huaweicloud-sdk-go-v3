package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportSubscriptionResponseReportContentSubscription **参数解释：** 报告内容订阅，配置订阅报告包含的具体统计模块启用状态。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportSubscriptionResponseReportContentSubscription struct {

	// **参数解释：** 是否启用总览统计内容（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	OverviewStatisticsEnable *bool `json:"overview_statistics_enable,omitempty"`

	// **参数解释：** 是否启用按天分组统计（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	GroupByDayEnable *bool `json:"group_by_day_enable,omitempty"`

	// **参数解释：** 是否启用请求次数统计内容（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	RequestStatisticsEnable *bool `json:"request_statistics_enable,omitempty"`

	// **参数解释：** 是否启用QPS统计内容（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	QpsStatisticsEnable *bool `json:"qps_statistics_enable,omitempty"`

	// **参数解释：** 是否启用带宽统计内容（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	BandwidthStatisticsEnable *bool `json:"bandwidth_statistics_enable,omitempty"`

	// **参数解释：** 是否启用响应码统计内容（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	ResponseCodeStatisticsEnable *bool `json:"response_code_statistics_enable,omitempty"`

	// **参数解释：** 是否启用攻击类型分布统计（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	AttackTypeDistributionEnable *bool `json:"attack_type_distribution_enable,omitempty"`

	// **参数解释：** 是否启用TOP被攻击域名统计（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	TopAttackedDomainsEnable *bool `json:"top_attacked_domains_enable,omitempty"`

	// **参数解释：** 是否启用TOP攻击源IP统计（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	TopAttackSourceIpsEnable *bool `json:"top_attack_source_ips_enable,omitempty"`

	// **参数解释：** 是否启用TOP被攻击URL统计（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	TopAttackedUrlsEnable *bool `json:"top_attacked_urls_enable,omitempty"`

	// **参数解释：** 是否启用TOP攻击源地理位置统计（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	TopAttackSourceLocationsEnable *bool `json:"top_attack_source_locations_enable,omitempty"`

	// **参数解释：** 是否启用TOP异常URL统计（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	TopAbnormalUrlsEnable *bool `json:"top_abnormal_urls_enable,omitempty"`
}

func (o SecurityReportSubscriptionResponseReportContentSubscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportSubscriptionResponseReportContentSubscription struct{}"
	}

	return strings.Join([]string{"SecurityReportSubscriptionResponseReportContentSubscription", string(data)}, " ")
}
