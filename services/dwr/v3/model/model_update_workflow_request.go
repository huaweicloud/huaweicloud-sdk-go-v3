package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowRequest Request Object
type UpdateWorkflowRequest struct {

	// 工作流名称。
	GraphName string `json:"graph_name"`

	Body *UpdateWorkflowBody `json:"body,omitempty"`
}

func (o UpdateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowRequest", string(data)}, " ")
}
