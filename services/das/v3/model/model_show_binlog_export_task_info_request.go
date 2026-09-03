package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBinlogExportTaskInfoRequest Request Object
type ShowBinlogExportTaskInfoRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 导出任务ID
	ExportTaskId int64 `json:"export_task_id"`
}

func (o ShowBinlogExportTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogExportTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBinlogExportTaskInfoRequest", string(data)}, " ")
}
