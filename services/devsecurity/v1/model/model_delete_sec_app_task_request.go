package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecAppTaskRequest Request Object
type DeleteSecAppTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteSecAppTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecAppTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecAppTaskRequest", string(data)}, " ")
}
