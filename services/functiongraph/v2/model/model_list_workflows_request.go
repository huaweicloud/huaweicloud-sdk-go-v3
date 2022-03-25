package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListWorkflowsRequest struct {
	// 函数流名称

	WorkflowName *string `json:"workflow_name,omitempty"`
	// 分页查询，每页显示的条目数量，最大数量200，超过200后只返回200

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询，分页的偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWorkflowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowsRequest", string(data)}, " ")
}
