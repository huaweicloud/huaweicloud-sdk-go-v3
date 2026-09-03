package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullSqlExportTasksRequest Request Object
type ListFullSqlExportTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 每页记录数
	PageSize int32 `json:"page_size"`

	// 页码
	PageNo int32 `json:"page_no"`
}

func (o ListFullSqlExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlExportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListFullSqlExportTasksRequest", string(data)}, " ")
}
