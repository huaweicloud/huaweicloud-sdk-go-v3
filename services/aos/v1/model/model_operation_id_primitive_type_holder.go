package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationIdPrimitiveTypeHolder struct {

	// 资源栈集操作Id。  此Id由资源编排服务在生成资源栈集操作的时候生成，为UUID。
	OperationId *string `json:"operation_id,omitempty"`
}

func (o OperationIdPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationIdPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"OperationIdPrimitiveTypeHolder", string(data)}, " ")
}
