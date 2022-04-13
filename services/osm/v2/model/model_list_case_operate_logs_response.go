package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseOperateLogsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseOperateLogsResponse", string(data)}, " ")
}
