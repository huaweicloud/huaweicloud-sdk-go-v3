package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadTaskReportRequest Request Object
type DownloadTaskReportRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DownloadTaskReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadTaskReportRequest struct{}"
	}

	return strings.Join([]string{"DownloadTaskReportRequest", string(data)}, " ")
}
