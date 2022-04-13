package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数流详情
type ListWorkflowsResult struct {
	// 唯一标识ID，流程定义ID

	Id *string `json:"id,omitempty"`
	// 唯一标识ID，流程URN

	WorkflowUrn *string `json:"workflow_urn,omitempty"`
	// 流程定义名称

	Name *string `json:"name,omitempty"`
	// 流程定义描述

	Description *string `json:"description,omitempty"`
	// 流程创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间

	CreatedTime *string `json:"created_time,omitempty"`
	// 流程修改时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间

	UpdatedTime *string `json:"updated_time,omitempty"`
	// 流程创建者

	CreatedBy *string `json:"created_by,omitempty"`
}

func (o ListWorkflowsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsResult struct{}"
	}

	return strings.Join([]string{"ListWorkflowsResult", string(data)}, " ")
}
