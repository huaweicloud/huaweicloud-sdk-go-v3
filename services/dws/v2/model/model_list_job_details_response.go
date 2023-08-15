package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobDetailsResponse Response Object
type ListJobDetailsResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`

	// 任务开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务当前状态
	Status *string `json:"status,omitempty"`

	// 任务失败错误码
	FailedCode *string `json:"failed_code,omitempty"`

	// 任务失败错误详情
	FailedDetail *string `json:"failed_detail,omitempty"`

	// 任务进度
	Progress       *string `json:"progress,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJobDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListJobDetailsResponse", string(data)}, " ")
}
