package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobResponse Response Object
type ShowJobResponse struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务类型
	JobType *string `json:"job_type,omitempty"`

	// 任务开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 任务错误码
	JobErrorCode *string `json:"job_error_code,omitempty"`

	// 任务失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 子任务总数
	SubJobsTotal *int32 `json:"sub_jobs_total,omitempty"`

	// 子任务列表
	SubJobs        *[]JobDetailInfo `json:"sub_jobs,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}
