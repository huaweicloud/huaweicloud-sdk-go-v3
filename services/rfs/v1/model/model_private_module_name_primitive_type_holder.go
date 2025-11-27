package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateModuleNamePrimitiveTypeHolder struct {

	// 私有模块（private-module）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ModuleName string `json:"module_name"`
}

func (o PrivateModuleNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateModuleNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivateModuleNamePrimitiveTypeHolder", string(data)}, " ")
}
