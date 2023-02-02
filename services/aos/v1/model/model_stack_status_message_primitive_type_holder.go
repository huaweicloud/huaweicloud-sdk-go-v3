package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackStatusMessagePrimitiveTypeHolder struct {

	// 当资源栈的状态为任意失败状态（即以 `FAILED` 结尾时），将会展示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o StackStatusMessagePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackStatusMessagePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackStatusMessagePrimitiveTypeHolder", string(data)}, " ")
}
