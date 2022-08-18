package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListWorkflowRequest struct {

	// 函数流名称
	WorkflowName *string `json:"workflow_name,omitempty"`

	// 分页查询，每页显示的条目数量，默认值为200 limit大于200时，按照200处理
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询，分页的偏移量，默认值为0 offset小于0时，按照0处理
	Offset *int32 `json:"offset,omitempty"`

	// 企业项目ID
	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	// 函数流模式 \"NORMAL\"标准函数流 \"EXPRESS\"快速函数流
	Mode *string `json:"mode,omitempty"`
}

func (o ListWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowRequest", string(data)}, " ")
}
