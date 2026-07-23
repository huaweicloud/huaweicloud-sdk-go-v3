package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobScheduleRequest Request Object
type CreateJobScheduleRequest struct {
	Body *CreateJobScheduleRequestBody `json:"body,omitempty"`
}

func (o CreateJobScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobScheduleRequest struct{}"
	}

	return strings.Join([]string{"CreateJobScheduleRequest", string(data)}, " ")
}
