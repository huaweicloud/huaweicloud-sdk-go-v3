package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateExportTaskResp struct {

	// 导出任务id
	TaskId int64 `json:"task_id"`
}

func (o CreateExportTaskResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExportTaskResp struct{}"
	}

	return strings.Join([]string{"CreateExportTaskResp", string(data)}, " ")
}
