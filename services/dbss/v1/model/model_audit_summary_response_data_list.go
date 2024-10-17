package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditSummaryResponseDataList struct {

	// ID
	Id *int64 `json:"id,omitempty"`

	// 状态 - 1: success - 2: failure
	Status *string `json:"status,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实例ID
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 审计时长
	AuditDuration *int64 `json:"audit_duration,omitempty"`

	// 语句总量
	TotalSql *int64 `json:"total_sql,omitempty"`

	// 风险总量
	TotalRisk *int64 `json:"total_risk,omitempty"`

	// 今日语句
	TodaySql *int64 `json:"today_sql,omitempty"`

	// 今日风险
	TodayRisk *int64 `json:"today_risk,omitempty"`

	// 今日会话
	TodaySession *int64 `json:"today_session,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 保留字1
	Reserve1 *string `json:"reserve1,omitempty"`

	// 保留字2
	Reserve2 *string `json:"reserve2,omitempty"`
}

func (o AuditSummaryResponseDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSummaryResponseDataList struct{}"
	}

	return strings.Join([]string{"AuditSummaryResponseDataList", string(data)}, " ")
}
