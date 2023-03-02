package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVideoSynthesisTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowVideoSynthesisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoSynthesisTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoSynthesisTaskRequest", string(data)}, " ")
}
