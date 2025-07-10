package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferWorkItemFlowRequest Request Object
type TransferWorkItemFlowRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *WorkItemFlowRequestBody `json:"body,omitempty"`
}

func (o TransferWorkItemFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferWorkItemFlowRequest struct{}"
	}

	return strings.Join([]string{"TransferWorkItemFlowRequest", string(data)}, " ")
}
