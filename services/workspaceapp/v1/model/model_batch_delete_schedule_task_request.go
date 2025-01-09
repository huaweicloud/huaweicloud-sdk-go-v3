package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScheduleTaskRequest Request Object
type BatchDeleteScheduleTaskRequest struct {
	Body *BatchDeleteScheduleTaskReq `json:"body,omitempty"`
}

func (o BatchDeleteScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScheduleTaskRequest", string(data)}, " ")
}
