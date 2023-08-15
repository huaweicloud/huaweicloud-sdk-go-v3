package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowRequest Request Object
type DeleteWorkflowRequest struct {

	// 工作流名称。名称必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于64个字符，且不能重名。
	GraphName string `json:"graph_name"`
}

func (o DeleteWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowRequest", string(data)}, " ")
}
