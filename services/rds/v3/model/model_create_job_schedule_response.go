package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobScheduleResponse Response Object
type CreateJobScheduleResponse struct {

	// 策略id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateJobScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobScheduleResponse struct{}"
	}

	return strings.Join([]string{"CreateJobScheduleResponse", string(data)}, " ")
}
