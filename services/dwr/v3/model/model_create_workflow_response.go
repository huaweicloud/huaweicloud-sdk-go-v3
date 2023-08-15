package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowResponse Response Object
type CreateWorkflowResponse struct {

	// 工作流名称。
	GraphName *string `json:"graph_name,omitempty"`

	// 工作流的URN。
	GraphUrn *string `json:"graph_urn,omitempty"`

	// 工作流创建的时间。
	CreatedAt *string `json:"created_at,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowResponse", string(data)}, " ")
}
