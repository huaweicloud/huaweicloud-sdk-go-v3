package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowExecutionsRequest Request Object
type ListWorkflowExecutionsRequest struct {

	// 任务id，从工作流命令列表中获取的工作流id。
	WorkflowId string `json:"workflow_id"`

	// 所属的企业项目id。
	XEnterpriseProjectId *string `json:"x_enterprise_project_id,omitempty"`
}

func (o ListWorkflowExecutionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowExecutionsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowExecutionsRequest", string(data)}, " ")
}
