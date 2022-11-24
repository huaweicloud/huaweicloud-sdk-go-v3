package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWorkflowRequest struct {

	// 工作流名称。名称必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于64个字符，且不能重名
	GraphName string `json:"graph_name"`

	Body *CreateWorkflowRequestBody `json:"body,omitempty"`
}

func (o CreateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowRequest", string(data)}, " ")
}
