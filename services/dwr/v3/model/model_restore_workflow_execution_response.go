package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreWorkflowExecutionResponse Response Object
type RestoreWorkflowExecutionResponse struct {

	// 运行实例的URN。
	ExecutionUrn *string `json:"execution_urn,omitempty"`

	// 运行实例的恢复启动时间。
	RestoredAt *string `json:"restored_at,omitempty"`

	// 运行实例的名字。
	ExecutionName *string `json:"execution_name,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreWorkflowExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreWorkflowExecutionResponse struct{}"
	}

	return strings.Join([]string{"RestoreWorkflowExecutionResponse", string(data)}, " ")
}
