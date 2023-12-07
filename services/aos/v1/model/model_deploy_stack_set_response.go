package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployStackSetResponse Response Object
type DeployStackSetResponse struct {

	// 资源栈集操作（stack_set_operation）的唯一Id。  此ID由资源编排服务在生成资源栈集操作的时候生成，为UUID。
	StackSetOperationId *string `json:"stack_set_operation_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o DeployStackSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployStackSetResponse struct{}"
	}

	return strings.Join([]string{"DeployStackSetResponse", string(data)}, " ")
}
