package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAopWorkflowVersionApprovelResponse Response Object
type CreateAopWorkflowVersionApprovelResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	Data           *WorkflowApproveOpinionDetail `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o CreateAopWorkflowVersionApprovelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAopWorkflowVersionApprovelResponse struct{}"
	}

	return strings.Join([]string{"CreateAopWorkflowVersionApprovelResponse", string(data)}, " ")
}
