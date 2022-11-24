package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackDescriptionPrimitiveTypeHolder struct {

	// 资源栈的描述。可用于客户识别自己的资源栈。
	Description *string `json:"description,omitempty"`
}

func (o StackDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
