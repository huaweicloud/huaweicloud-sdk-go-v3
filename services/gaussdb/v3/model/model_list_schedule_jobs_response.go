package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListScheduleJobsResponse struct {

	// 任务详情
	Schedules *[]ScheduleTask `json:"schedules,omitempty"`

	// 记录总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScheduleJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleJobsResponse struct{}"
	}

	return strings.Join([]string{"ListScheduleJobsResponse", string(data)}, " ")
}
