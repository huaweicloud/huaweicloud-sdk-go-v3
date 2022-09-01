package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNewTaskStatusRequest struct {

	// 任务ID（非负整数）的字符串
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`
}

func (o ShowNewTaskStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNewTaskStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowNewTaskStatusRequest", string(data)}, " ")
}
