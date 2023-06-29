package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowResponse Response Object
type UpdateWorkflowResponse struct {

	// 工作流名称。
	GraphName *string `json:"graph_name,omitempty"`

	// 工作流的URN。
	GraphUrn *string `json:"graph_urn,omitempty"`

	// 工作流更新的时间。
	LastModified *string `json:"last_modified,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowResponse", string(data)}, " ")
}
