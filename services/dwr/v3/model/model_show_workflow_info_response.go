package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowInfoResponse Response Object
type ShowWorkflowInfoResponse struct {

	// 工作流的名称。
	Name *string `json:"name,omitempty"`

	// 工作流的创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 工作流的定义。
	Definition *interface{} `json:"definition,omitempty"`

	// 工作流的URN。
	GraphUrn *string `json:"graph_urn,omitempty"`

	// 工作流的描述。
	Description *string `json:"description,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWorkflowInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowInfoResponse", string(data)}, " ")
}
