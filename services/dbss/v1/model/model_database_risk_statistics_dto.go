package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseRiskStatisticsDto struct {

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 高风险总量
	HighRisk *int64 `json:"high_risk,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 低风险总量
	LowRisk *int64 `json:"low_risk,omitempty"`

	// 中风险总量
	MediumRisk *int64 `json:"medium_risk,omitempty"`

	// 风险规则统计
	RiskRule *[]RiskRuleStatistic `json:"risk_rule,omitempty"`

	// 总风险数量
	TotalRisk *int64 `json:"total_risk,omitempty"`
}

func (o DatabaseRiskStatisticsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseRiskStatisticsDto struct{}"
	}

	return strings.Join([]string{"DatabaseRiskStatisticsDto", string(data)}, " ")
}
