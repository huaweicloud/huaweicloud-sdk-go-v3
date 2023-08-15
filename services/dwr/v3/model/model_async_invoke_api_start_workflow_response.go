package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsyncInvokeApiStartWorkflowResponse Response Object
type AsyncInvokeApiStartWorkflowResponse struct {

	// 运行实例的URN。
	ExecutionUrn *string `json:"execution_urn,omitempty"`

	// 运行实例启动时间。
	StartedAt *string `json:"started_at,omitempty"`

	// 运行实例的名字。
	ExecutionName *string `json:"execution_name,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AsyncInvokeApiStartWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncInvokeApiStartWorkflowResponse struct{}"
	}

	return strings.Join([]string{"AsyncInvokeApiStartWorkflowResponse", string(data)}, " ")
}
