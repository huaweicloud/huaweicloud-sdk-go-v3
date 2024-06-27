package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelCaptureTaskDto struct {

	// 抓包任务id
	TaskId *string `json:"task_id,omitempty"`
}

func (o CancelCaptureTaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelCaptureTaskDto struct{}"
	}

	return strings.Join([]string{"CancelCaptureTaskDto", string(data)}, " ")
}
