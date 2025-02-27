package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleRecordTasksRequest Request Object
type DeleteScheduleRecordTasksRequest struct {

	// 已存在的计划录制任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteScheduleRecordTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleRecordTasksRequest struct{}"
	}

	return strings.Join([]string{"DeleteScheduleRecordTasksRequest", string(data)}, " ")
}
