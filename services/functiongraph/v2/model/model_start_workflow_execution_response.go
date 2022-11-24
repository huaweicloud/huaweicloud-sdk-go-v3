package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartWorkflowExecutionResponse struct {

	// 错误码
	ExecutionId    *string `json:"execution_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartWorkflowExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartWorkflowExecutionResponse struct{}"
	}

	return strings.Join([]string{"StartWorkflowExecutionResponse", string(data)}, " ")
}
