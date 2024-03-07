package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRecordsDetailsRequest Request Object
type ListScheduledTasksRecordsDetailsRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`

	// 任务执行记录ID。
	RecordId string `json:"record_id"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset int32 `json:"offset"`

	// 用于分页查询，每页返回的个数，取值范围0~50。
	Limit int32 `json:"limit"`
}

func (o ListScheduledTasksRecordsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRecordsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRecordsDetailsRequest", string(data)}, " ")
}
