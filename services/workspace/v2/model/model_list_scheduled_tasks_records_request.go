package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRecordsRequest Request Object
type ListScheduledTasksRecordsRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。取值范围0-100，默认值是10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduledTasksRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRecordsRequest", string(data)}, " ")
}
