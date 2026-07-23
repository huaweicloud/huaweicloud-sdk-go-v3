package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobScheduleRequest Request Object
type DeleteJobScheduleRequest struct {

	// 策略ID
	ScheduleId string `json:"schedule_id"`
}

func (o DeleteJobScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobScheduleRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobScheduleRequest", string(data)}, " ")
}
