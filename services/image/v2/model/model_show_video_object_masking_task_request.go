package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVideoObjectMaskingTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowVideoObjectMaskingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoObjectMaskingTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoObjectMaskingTaskRequest", string(data)}, " ")
}
