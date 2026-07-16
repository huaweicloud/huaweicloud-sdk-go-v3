package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowRequest Request Object
type ShowWorkflowRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ShowWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowRequest", string(data)}, " ")
}
