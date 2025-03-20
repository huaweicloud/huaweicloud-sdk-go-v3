package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduledTaskRequest Request Object
type CreateScheduledTaskRequest struct {
	Body *ScheduledTaskRequestBody `json:"body,omitempty"`
}

func (o CreateScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateScheduledTaskRequest", string(data)}, " ")
}
