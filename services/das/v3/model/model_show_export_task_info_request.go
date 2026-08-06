package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExportTaskInfoRequest Request Object
type ShowExportTaskInfoRequest struct {

	// 任务ID
	TaskId float32 `json:"task_id"`
}

func (o ShowExportTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExportTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowExportTaskInfoRequest", string(data)}, " ")
}
