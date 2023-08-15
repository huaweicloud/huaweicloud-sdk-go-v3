package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageHighresolutionMattingTaskRequest Request Object
type ShowImageHighresolutionMattingTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowImageHighresolutionMattingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageHighresolutionMattingTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowImageHighresolutionMattingTaskRequest", string(data)}, " ")
}
