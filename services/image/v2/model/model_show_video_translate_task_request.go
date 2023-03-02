package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVideoTranslateTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowVideoTranslateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoTranslateTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoTranslateTaskRequest", string(data)}, " ")
}
