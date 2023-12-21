package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobDetailInfo struct {

	// 子任务ID。
	Id *string `json:"id,omitempty"`

	JobType *JobType `json:"job_type,omitempty"`

	JobResourceInfo *JobResourceInfo `json:"job_resource_info,omitempty"`

	// 任务创建时间。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 任务结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Status *JobDetailStatus `json:"status,omitempty"`

	// 任务执行失败时的错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 任务失败原因。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o JobDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetailInfo struct{}"
	}

	return strings.Join([]string{"JobDetailInfo", string(data)}, " ")
}
