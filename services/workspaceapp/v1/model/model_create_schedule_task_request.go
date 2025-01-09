package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleTaskRequest Request Object
type CreateScheduleTaskRequest struct {
	Body *CreateScheduleTaskReq `json:"body,omitempty"`
}

func (o CreateScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateScheduleTaskRequest", string(data)}, " ")
}
