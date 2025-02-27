package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleRecordTasksRequest Request Object
type CreateScheduleRecordTasksRequest struct {
	Body *ScheduleRecordTasksReq `json:"body,omitempty"`
}

func (o CreateScheduleRecordTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleRecordTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateScheduleRecordTasksRequest", string(data)}, " ")
}
