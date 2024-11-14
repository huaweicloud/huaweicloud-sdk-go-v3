package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScheduledTaskRequest Request Object
type CancelScheduledTaskRequest struct {

	// 任务ID，取值为定时任务列表。
	JobId string `json:"job_id"`
}

func (o CancelScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelScheduledTaskRequest", string(data)}, " ")
}
