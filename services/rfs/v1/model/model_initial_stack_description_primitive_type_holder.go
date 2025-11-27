package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InitialStackDescriptionPrimitiveTypeHolder struct {

	// 初始化资源栈描述。可用于客户识别被资源栈集所管理的资源栈。  资源栈集下的资源栈仅在创建时统一使用该描述。客户想要更新初始化资源栈描述，可以通过UpdateStackSet API。  后续更新资源栈集描述将不会同步更新已管理的资源栈描述。
	InitialStackDescription *string `json:"initial_stack_description,omitempty"`
}

func (o InitialStackDescriptionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitialStackDescriptionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"InitialStackDescriptionPrimitiveTypeHolder", string(data)}, " ")
}
