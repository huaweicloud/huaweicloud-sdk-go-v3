package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityReportSubscriptionRequestBodyReportContentSubscription **参数解释：** 报告内容订阅 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpdateSecurityReportSubscriptionRequestBodyReportContentSubscription struct {
	OverviewStatisticsEnable *bool `json:"overview_statistics_enable,omitempty"`

	GroupByDayEnable *bool `json:"group_by_day_enable,omitempty"`

	RequestStatisticsEnable *bool `json:"request_statistics_enable,omitempty"`

	QpsStatisticsEnable *bool `json:"qps_statistics_enable,omitempty"`

	BandwidthStatisticsEnable *bool `json:"bandwidth_statistics_enable,omitempty"`

	ResponseCodeStatisticsEnable *bool `json:"response_code_statistics_enable,omitempty"`

	AttackTypeDistributionEnable *bool `json:"attack_type_distribution_enable,omitempty"`

	TopAttackedDomainsEnable *bool `json:"top_attacked_domains_enable,omitempty"`

	TopAttackSourceIpsEnable *bool `json:"top_attack_source_ips_enable,omitempty"`

	TopAttackedUrlsEnable *bool `json:"top_attacked_urls_enable,omitempty"`

	TopAttackSourceLocationsEnable *bool `json:"top_attack_source_locations_enable,omitempty"`

	TopAbnormalUrlsEnable *bool `json:"top_abnormal_urls_enable,omitempty"`
}

func (o UpdateSecurityReportSubscriptionRequestBodyReportContentSubscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityReportSubscriptionRequestBodyReportContentSubscription struct{}"
	}

	return strings.Join([]string{"UpdateSecurityReportSubscriptionRequestBodyReportContentSubscription", string(data)}, " ")
}
