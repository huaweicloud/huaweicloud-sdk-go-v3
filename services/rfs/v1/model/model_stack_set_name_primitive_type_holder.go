package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetNamePrimitiveTypeHolder struct {

	// 资源栈集（stack_set）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`
}

func (o StackSetNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetNamePrimitiveTypeHolder", string(data)}, " ")
}
