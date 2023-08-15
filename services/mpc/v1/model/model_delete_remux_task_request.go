package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRemuxTaskRequest Request Object
type DeleteRemuxTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteRemuxTaskRequest", string(data)}, " ")
}
