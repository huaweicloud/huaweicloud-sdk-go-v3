package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto update body Object
type ModifyInstanceNameReq struct {
	// DDM实例名称，命名要求如下。 - 长度为4-64个字符。 - 必须以字母开头。 - 可以包含字母、数字、中划线，不能包含其它特殊字符。

	Name string `json:"name"`
}

func (o ModifyInstanceNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceNameReq struct{}"
	}

	return strings.Join([]string{"ModifyInstanceNameReq", string(data)}, " ")
}
