package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OverviewSubscriptionRes 安全总览订阅的项目
type OverviewSubscriptionRes struct {

	// **参数解释**: 安全总览 **取值范围**: - true：开启 - false：关闭
	DashboardStatus *bool `json:"dashboard_status,omitempty"`

	// **参数解释**: 风险趋势 **取值范围**: - true：开启 - false：关闭
	RiskyTrendStatus *bool `json:"risky_trend_status,omitempty"`

	// **参数解释**: 风险分布 **取值范围**: - true：开启 - false：关闭
	RiskyDistributionStatus *bool `json:"risky_distribution_status,omitempty"`

	// **参数解释**: top5风险主机 **取值范围**: - true：开启 - false：关闭
	TopRiskyHostStatus *bool `json:"top_risky_host_status,omitempty"`

	// **参数解释**: top5暴力破解来源 **取值范围**: - true：开启 - false：关闭
	TopAccountCrackStatus *bool `json:"top_account_crack_status,omitempty"`

	// **参数解释**: 服务器列表来源 **取值范围**: - true：开启 - false：关闭
	HostInfoStatus *bool `json:"host_info_status,omitempty"`
}

func (o OverviewSubscriptionRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverviewSubscriptionRes struct{}"
	}

	return strings.Join([]string{"OverviewSubscriptionRes", string(data)}, " ")
}
