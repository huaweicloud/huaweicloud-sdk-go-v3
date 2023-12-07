package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetOperationIdPrimitiveTypeHolder struct {

	// 资源栈集操作（stack_set_operation）的唯一Id。  此ID由资源编排服务在生成资源栈集操作的时候生成，为UUID。
	StackSetOperationId *string `json:"stack_set_operation_id,omitempty"`
}

func (o StackSetOperationIdPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetOperationIdPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetOperationIdPrimitiveTypeHolder", string(data)}, " ")
}
