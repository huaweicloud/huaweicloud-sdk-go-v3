package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduledTasksRequest Request Object
type CreateScheduledTasksRequest struct {
	Body *CreateScheduledTasksReq `json:"body,omitempty"`
}

func (o CreateScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateScheduledTasksRequest", string(data)}, " ")
}
