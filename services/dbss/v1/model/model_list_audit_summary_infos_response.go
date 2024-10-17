package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSummaryInfosResponse Response Object
type ListAuditSummaryInfosResponse struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 审计总时长
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

	// 列表信息
	DataList *[]AuditSummaryResponseDataList `json:"data_list,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditSummaryInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSummaryInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAuditSummaryInfosResponse", string(data)}, " ")
}
