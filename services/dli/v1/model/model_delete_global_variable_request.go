package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalVariableRequest Request Object
type DeleteGlobalVariableRequest struct {

	// 全局变量名，名称只能包含数字、英文字母和下划线，但不能是纯数字，不能以下划线开头，且不能超过128字符
	VarName string `json:"var_name"`
}

func (o DeleteGlobalVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalVariableRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalVariableRequest", string(data)}, " ")
}
