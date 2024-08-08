package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageJobDetailInfo struct {

	// 子任务ID。
	Id *string `json:"id,omitempty"`

	JobType *ImageJobType `json:"job_type,omitempty"`

	JobResourceInfo *ImageJobResourceInfo `json:"job_resource_info,omitempty"`

	// 任务创建时间。
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 任务结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Status *ImageJobDetailStatus `json:"status,omitempty"`

	JobExecuteInfo *ImageJobExecuteInfo `json:"job_execute_info,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ImageJobDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageJobDetailInfo struct{}"
	}

	return strings.Join([]string{"ImageJobDetailInfo", string(data)}, " ")
}
