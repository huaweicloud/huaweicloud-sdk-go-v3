package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRunResponse struct {

	// 作业运行ID。
	RunId *string `json:"run_id,omitempty"`

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业类型。
	JobType *string `json:"job_type,omitempty"`

	// 此作业的当前状态，包含提交（LAUNCHING）、运行中（RUNNING）、完成（FINISHED）、失败（FAILED）、取消（CANCELLED）。
	Status *string `json:"status,omitempty"`

	// 创建运行时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// 系统提示信息。运行失败时，失败原因。
	Message *string `json:"message,omitempty"`

	// 作业运行详情。
	Details        *[]RunDetail `json:"details,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRunResponse struct{}"
	}

	return strings.Join([]string{"ShowRunResponse", string(data)}, " ")
}
