package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateScheduleRequestBody struct {
	Schedule *ScheduleObj `json:"schedule"`

	Parameters *GcParameters `json:"parameters,omitempty"`
}

func (o UpdateScheduleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateScheduleRequestBody", string(data)}, " ")
}
