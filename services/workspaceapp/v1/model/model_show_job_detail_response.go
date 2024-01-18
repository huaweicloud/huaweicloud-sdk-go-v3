package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobDetailResponse Response Object
type ShowJobDetailResponse struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	JobType *JobType `json:"job_type,omitempty"`

	// 任务创建时间。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 任务结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	// 子任务总数。
	SubJobsTotal *int32 `json:"sub_jobs_total,omitempty"`

	// 子任务列表。
	SubJobs        *[]JobDetailInfo `json:"sub_jobs,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
