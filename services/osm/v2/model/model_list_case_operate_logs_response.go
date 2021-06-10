package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCaseOperateLogsResponse struct {
	// 总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 工单操作日志列表

	IncidentOperateLogList *[]IncidentOperateLogV2 `json:"incident_operate_log_list,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o ListCaseOperateLogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCaseOperateLogsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseOperateLogsResponse", string(data)}, " ")
}
