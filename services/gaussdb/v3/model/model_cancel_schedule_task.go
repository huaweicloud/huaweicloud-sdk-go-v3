package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScheduleTask 取消定时任务请求体
type CancelScheduleTask struct {

	// 任务ID。
	JobIds []string `json:"job_ids"`
}

func (o CancelScheduleTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScheduleTask struct{}"
	}

	return strings.Join([]string{"CancelScheduleTask", string(data)}, " ")
}
