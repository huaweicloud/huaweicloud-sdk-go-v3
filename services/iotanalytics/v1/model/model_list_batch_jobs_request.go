package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBatchJobsRequest Request Object
type ListBatchJobsRequest struct {

	// 当前偏移量，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的最大作业个数，范围: [1, 100]。默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 是否定时作业。true：定时作业：false：不是定时作业。为空时：所有作业。
	HasSchedule *bool `json:"has_schedule,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 调度状态。1:NORMAL, 2:PAUSED, 3:COMPLETE, 4:ERROR, 5:BLOCKED
	ScheduleStatus *string `json:"schedule_status,omitempty"`

	// 指定作业排序字段，默认为created_time（作业创建时间），支持created_time(作业创建时间)、modified_time（作业更新时间） 、job_name（作业名称）三种排序字段。
	OrderBy *string `json:"order_by,omitempty"`

	// 指定作业排序的升降序，默认为desc（降序），支持asc（升序）、desc（降序）两种排序方式。
	Order *string `json:"order,omitempty"`
}

func (o ListBatchJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBatchJobsRequest struct{}"
	}

	return strings.Join([]string{"ListBatchJobsRequest", string(data)}, " ")
}
