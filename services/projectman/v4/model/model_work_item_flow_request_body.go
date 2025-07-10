package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowRequestBody 工作项流程流转请求参数
type WorkItemFlowRequestBody struct {

	// 工作项唯一Id
	Id *string `json:"id,omitempty"`

	// 工作项类型
	IssueCategory *string `json:"issue_category,omitempty"`

	// 工作项流转code
	FlowCode *string `json:"flow_code,omitempty"`
}

func (o WorkItemFlowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowRequestBody struct{}"
	}

	return strings.Join([]string{"WorkItemFlowRequestBody", string(data)}, " ")
}
