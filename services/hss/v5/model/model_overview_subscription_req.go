package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OverviewSubscriptionReq 安全总览订阅信息
type OverviewSubscriptionReq struct {

	// **参数解释**： 安全总览是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	DashboardStatus *bool `json:"dashboard_status,omitempty"`

	// **参数解释**： 风险趋势是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	RiskyTrendStatus *bool `json:"risky_trend_status,omitempty"`

	// **参数解释**： 风险分布是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	RiskyDistributionStatus *bool `json:"risky_distribution_status,omitempty"`

	// **参数解释**： top5风险主机是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	TopRiskyHostStatus *bool `json:"top_risky_host_status,omitempty"`

	// **参数解释**： top5暴力破解来源是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	TopAccountCrackStatus *bool `json:"top_account_crack_status,omitempty"`

	// **参数解释**： 服务器列表是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	HostInfoStatus *bool `json:"host_info_status,omitempty"`
}

func (o OverviewSubscriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverviewSubscriptionReq struct{}"
	}

	return strings.Join([]string{"OverviewSubscriptionReq", string(data)}, " ")
}
