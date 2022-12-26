package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteWorkflowResponse struct {

	// 执行ID
	ExecutionId    *string `json:"execution_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteWorkflowResponse struct{}"
	}

	return strings.Join([]string{"ExecuteWorkflowResponse", string(data)}, " ")
}
