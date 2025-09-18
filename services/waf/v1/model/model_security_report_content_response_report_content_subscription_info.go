package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportContentResponseReportContentSubscriptionInfo **参数解释：** 报告内容订阅，包含安全报告的各类统计数据详情。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportContentResponseReportContentSubscriptionInfo struct {

	// **参数解释：** 总览统计信息，包含各维度的汇总统计数据及TOP域名详情。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	OverviewStatisticsListInfo *[]SecurityReportContentResponseReportContentSubscriptionInfoOverviewStatisticsListInfo `json:"overview_statistics_list_info,omitempty"`

	// **参数解释：** 请求次数统计信息，包含各维度按时间线的请求数量统计。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RequestStatisticsInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoRequestStatisticsInfoList `json:"request_statistics_info_list,omitempty"`

	QpsStatisticsInfo *SecurityReportContentResponseReportContentSubscriptionInfoQpsStatisticsInfo `json:"qps_statistics_info,omitempty"`

	BandwidthStatisticsInfo *SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfo `json:"bandwidth_statistics_info,omitempty"`

	ResponseCodeStatisticsInfo *SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfo `json:"response_code_statistics_info,omitempty"`

	// **参数解释：** 攻击类型分布统计信息，包含各攻击类型的数量分布。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AttackTypeDistributionInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoAttackTypeDistributionInfoList `json:"attack_type_distribution_info_list,omitempty"`

	// **参数解释：** TOP被攻击的域名信息，按被攻击次数排序的域名列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopAttackedDomainsInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoTopAttackedDomainsInfoList `json:"top_attacked_domains_info_list,omitempty"`

	// **参数解释：** TOP攻击的源IP信息，按攻击次数排序的攻击源IP列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopAttackSourceIpsInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoTopAttackSourceIpsInfoList `json:"top_attack_source_ips_info_list,omitempty"`

	// **参数解释：** TOP被攻击的URL信息，按被攻击次数排序的URL列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopAttackedUrlsInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoTopAttackedUrlsInfoList `json:"top_attacked_urls_info_list,omitempty"`

	// **参数解释：** TOP攻击的源地理位置信息，按攻击次数排序的地理位置列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopAttackSourceLocationsInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoTopAttackSourceLocationsInfoList `json:"top_attack_source_locations_info_list,omitempty"`

	TopAbnormalUrlsInfo *SecurityReportContentResponseReportContentSubscriptionInfoTopAbnormalUrlsInfo `json:"top_abnormal_urls_info,omitempty"`
}

func (o SecurityReportContentResponseReportContentSubscriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentSubscriptionInfo struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentSubscriptionInfo", string(data)}, " ")
}
