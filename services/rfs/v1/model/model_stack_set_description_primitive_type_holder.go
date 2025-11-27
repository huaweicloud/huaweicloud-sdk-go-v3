package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetDescriptionPrimitiveTypeHolder struct {

	// 资源栈集的描述。可用于客户识别自己的资源栈集。
	StackSetDescription *string `json:"stack_set_description,omitempty"`
}

func (o StackSetDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
