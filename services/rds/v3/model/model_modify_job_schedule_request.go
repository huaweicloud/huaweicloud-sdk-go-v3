package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyJobScheduleRequest Request Object
type ModifyJobScheduleRequest struct {

	// 策略ID
	ScheduleId string `json:"schedule_id"`

	Body *ModifyJobScheduleRequestBody `json:"body,omitempty"`
}

func (o ModifyJobScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobScheduleRequest struct{}"
	}

	return strings.Join([]string{"ModifyJobScheduleRequest", string(data)}, " ")
}
