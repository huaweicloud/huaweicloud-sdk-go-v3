package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportImageSecurityReportTaskResponse Response Object
type ExportImageSecurityReportTaskResponse struct {

	// 导出任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportImageSecurityReportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageSecurityReportTaskResponse struct{}"
	}

	return strings.Join([]string{"ExportImageSecurityReportTaskResponse", string(data)}, " ")
}
