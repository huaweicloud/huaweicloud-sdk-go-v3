package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionLabelsRequest Request Object
type ListExecutionLabelsRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`
}

func (o ListExecutionLabelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionLabelsRequest struct{}"
	}

	return strings.Join([]string{"ListExecutionLabelsRequest", string(data)}, " ")
}
