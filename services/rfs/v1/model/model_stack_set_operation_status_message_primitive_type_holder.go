package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetOperationStatusMessagePrimitiveTypeHolder struct {

	// 资源栈集操作失败时会展示此次操作失败的原因，例如，资源栈实例部署或删除失败个数超过上限或资源栈集操作超时。  如果需要查看详细失败信息，可通过ListStackInstances API获取查看资源栈实例的status_message。
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o StackSetOperationStatusMessagePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetOperationStatusMessagePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetOperationStatusMessagePrimitiveTypeHolder", string(data)}, " ")
}
