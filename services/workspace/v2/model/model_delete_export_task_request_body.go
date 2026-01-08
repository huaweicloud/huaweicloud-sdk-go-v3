package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTaskRequestBody 删除导出任务。
type DeleteExportTaskRequestBody struct {

	// 删除的导出任务id。
	TaskIds *[]string `json:"task_ids,omitempty"`
}

func (o DeleteExportTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTaskRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteExportTaskRequestBody", string(data)}, " ")
}
