package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageJobInfo struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	JobType *ImageJobType `json:"job_type,omitempty"`

	// 任务创建时间。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 任务结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Status *ImageJobStatus `json:"status,omitempty"`

	// 子任务总数。
	SubJobsTotal *int32 `json:"sub_jobs_total,omitempty"`
}

func (o ImageJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageJobInfo struct{}"
	}

	return strings.Join([]string{"ImageJobInfo", string(data)}, " ")
}
