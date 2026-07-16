package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowPurchasePoolRequest Request Object
type CreateWorkflowPurchasePoolRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	Body *WorkflowServicePackage `json:"body,omitempty"`
}

func (o CreateWorkflowPurchasePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowPurchasePoolRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowPurchasePoolRequest", string(data)}, " ")
}
