package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportScheduledTasksRecordsRequest Request Object
type ExportScheduledTasksRecordsRequest struct {

	// 任务ID。
	TaskId string `json:"task_id"`

	// 语言。
	Language *string `json:"language,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ExportScheduledTasksRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportScheduledTasksRecordsRequest struct{}"
	}

	return strings.Join([]string{"ExportScheduledTasksRecordsRequest", string(data)}, " ")
}
