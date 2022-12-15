package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubJobs struct {

	// job id
	JobId *string `json:"job_id,omitempty"`

	// job类型
	JobType *string `json:"job_type,omitempty"`

	// 创建时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 创建完成时间
	EndTime *string `json:"end_time,omitempty"`

	// job状态
	Status *string `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	FailReason *string `json:"fail_reason,omitempty"`

	Entities *SubJobs `json:"entities,omitempty"`
}

func (o SubJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobs struct{}"
	}

	return strings.Join([]string{"SubJobs", string(data)}, " ")
}
