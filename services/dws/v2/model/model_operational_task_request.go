package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationalTaskRequest struct {
	TaskInfo *TaskInfo `json:"task_info,omitempty"`
}

func (o OperationalTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationalTaskRequest struct{}"
	}

	return strings.Join([]string{"OperationalTaskRequest", string(data)}, " ")
}
