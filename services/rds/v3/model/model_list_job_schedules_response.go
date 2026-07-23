package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobSchedulesResponse Response Object
type ListJobSchedulesResponse struct {

	// 策略列表。
	Schedules *[]JobScheduleInfo `json:"schedules,omitempty"`

	// 策略总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListJobSchedulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobSchedulesResponse struct{}"
	}

	return strings.Join([]string{"ListJobSchedulesResponse", string(data)}, " ")
}
