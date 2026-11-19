package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OverviewInfoRes 安全总览信息
type OverviewInfoRes struct {

	// **参数解释**： 安全总览内容列表 **取值范围**： 0-100个DashboardInfo对象
	DashboardInfoList *[]interface{} `json:"dashboard_info_list,omitempty"`

	// **参数解释**： 风险趋势内容列表 **取值范围**： 0-100个RiskyTrendInfo对象
	RiskyTrendInfoList *[]interface{} `json:"risky_trend_info_list,omitempty"`

	// **参数解释**： 风险分布内容列表 **取值范围**： 0-100个RiskyDistributionInfo对象
	RiskyDistributionInfoList *[]interface{} `json:"risky_distribution_info_list,omitempty"`

	// **参数解释**： top5风险主机列表 **取值范围**： 0-5个TopRiskyHostInfo对象
	TopRiskyHostInfoList *[]interface{} `json:"top_risky_host_info_list,omitempty"`

	// **参数解释**： top5暴力破解来源列表 **取值范围**： 0-5个TopAccountCrackInfo对象
	TopAccountCrackInfoList *[]interface{} `json:"top_account_crack_info_list,omitempty"`

	// **参数解释**： 服务器列表 **取值范围**： 0-200个HostInfo对象
	HostInfoList *[]interface{} `json:"host_info_list,omitempty"`

	// **参数解释**： 总体情况列表 **取值范围**： 0-100个OverallSituationInfo对象
	OverallSituationInfoList *[]interface{} `json:"overall_situation_info_list,omitempty"`
}

func (o OverviewInfoRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverviewInfoRes struct{}"
	}

	return strings.Join([]string{"OverviewInfoRes", string(data)}, " ")
}
