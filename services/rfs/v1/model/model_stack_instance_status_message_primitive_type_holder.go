package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackInstanceStatusMessagePrimitiveTypeHolder struct {

	// 在资源栈实例状态为`INOPERABLE`或`OPERATION_FAILED`时，会显示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o StackInstanceStatusMessagePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackInstanceStatusMessagePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackInstanceStatusMessagePrimitiveTypeHolder", string(data)}, " ")
}
