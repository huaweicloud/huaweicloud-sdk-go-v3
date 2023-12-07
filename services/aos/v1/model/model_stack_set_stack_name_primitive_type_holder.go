package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetStackNamePrimitiveTypeHolder struct {

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName *string `json:"stack_name,omitempty"`
}

func (o StackSetStackNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetStackNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetStackNamePrimitiveTypeHolder", string(data)}, " ")
}
