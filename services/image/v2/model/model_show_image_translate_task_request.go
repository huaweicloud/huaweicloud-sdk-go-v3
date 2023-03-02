package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowImageTranslateTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowImageTranslateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageTranslateTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowImageTranslateTaskRequest", string(data)}, " ")
}
