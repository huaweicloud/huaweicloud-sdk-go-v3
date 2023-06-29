package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowTriggersRequest Request Object
type ListWorkflowTriggersRequest struct {

	// 目标函数流的id
	WorkflowId string `json:"workflow_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`
}

func (o ListWorkflowTriggersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowTriggersRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowTriggersRequest", string(data)}, " ")
}
