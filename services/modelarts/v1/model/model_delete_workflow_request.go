package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowRequest Request Object
type DeleteWorkflowRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`
}

func (o DeleteWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowRequest", string(data)}, " ")
}
