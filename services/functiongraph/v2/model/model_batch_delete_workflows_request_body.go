package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数流批量删除Body体
type BatchDeleteWorkflowsRequestBody struct {
	// 流程URN列表

	WorkflowUrns *[]string `json:"workflow_urns,omitempty"`
}

func (o BatchDeleteWorkflowsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkflowsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkflowsRequestBody", string(data)}, " ")
}
