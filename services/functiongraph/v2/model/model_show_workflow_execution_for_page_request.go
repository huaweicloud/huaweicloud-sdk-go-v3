package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWorkflowExecutionForPageRequest struct {

	// 函数工作流ID
	WorkflowId string `json:"workflow_id"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0
	Offset int32 `json:"offset"`

	// 分页查询，每页查询数据条数，取值范围：1,2,3...100
	Limit int32 `json:"limit"`

	// 查询开始时间，UTC时间，格式：YYYY-MM-DD hh:mm:ss。若起始时间未填写，以终止时间前推3天为起始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间，UTC时间，格式：YYYY-MM-DD hh:mm:ss。若终止时间未填写，以起始时间后退3天未终止时间。若均未填写，默认查询最近3天数据。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowWorkflowExecutionForPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowExecutionForPageRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowExecutionForPageRequest", string(data)}, " ")
}
