package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowStatisticResponse Response Object
type ListWorkflowStatisticResponse struct {

	// 应用总数
	AppCount *int32 `json:"app_count,omitempty"`

	// 流程总数
	WorkflowCount *int32 `json:"workflow_count,omitempty"`

	// 作业总数
	JobCount *int32 `json:"job_count,omitempty"`

	// 运行成功作业总数
	SucceedJobCount *int32 `json:"succeed_job_count,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ListWorkflowStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowStatisticResponse", string(data)}, " ")
}
