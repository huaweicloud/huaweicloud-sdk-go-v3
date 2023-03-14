package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTextToImageTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowTextToImageTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTextToImageTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTextToImageTaskRequest", string(data)}, " ")
}
