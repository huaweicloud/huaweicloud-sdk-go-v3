package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditStatisticsResponse Response Object
type ShowAuditStatisticsResponse struct {

	// 数据库风险统计
	DbRiskStatistics *[]DatabaseRiskStatisticsDto `json:"db_risk_statistics,omitempty"`

	// 高风险总量
	HighRiskTotal *int64 `json:"high_risk_total,omitempty"`

	// 低风险总量
	LowRiskTotal *int64 `json:"low_risk_total,omitempty"`

	// 中风险总量
	MediumRiskTotal *int64 `json:"medium_risk_total,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 风险规则统计
	RuleRiskStatistics *[]RuleRiskStatisticsDto `json:"rule_risk_statistics,omitempty"`

	UnsupportedAuditInfos *UnsupportedAuditInfoResponse `json:"unsupported_audit_infos,omitempty"`

	// 更新时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAuditStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditStatisticsResponse", string(data)}, " ")
}
