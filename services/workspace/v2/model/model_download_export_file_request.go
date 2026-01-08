package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadExportFileRequest Request Object
type DownloadExportFileRequest struct {

	// 导出任务id。
	TaskId string `json:"task_id"`
}

func (o DownloadExportFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadExportFileRequest struct{}"
	}

	return strings.Join([]string{"DownloadExportFileRequest", string(data)}, " ")
}
