package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfo **参数解释：** 响应码统计信息，包含WAF和上游服务器各响应码的时间线统计。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfo struct {

	// **参数解释：** WAF响应码统计列表，包含各响应码按时间线的WAF返回数量。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ResponseSourceWafInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfoResponseSourceWafInfoList `json:"response_source_waf_info_list,omitempty"`

	// **参数解释：** 上游响应码统计列表，包含各响应码按时间线的上游返回数量。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ResponseSourceUpstreamInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfoResponseSourceUpstreamInfoList `json:"response_source_upstream_info_list,omitempty"`
}

func (o SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfo struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfo", string(data)}, " ")
}
