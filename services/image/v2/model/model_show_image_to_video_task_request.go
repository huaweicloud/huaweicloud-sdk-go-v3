package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowImageToVideoTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowImageToVideoTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageToVideoTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowImageToVideoTaskRequest", string(data)}, " ")
}
