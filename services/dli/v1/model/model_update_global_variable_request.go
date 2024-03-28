package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalVariableRequest Request Object
type UpdateGlobalVariableRequest struct {

	// 全局变量名，名称只能包含数字、英文字母和下划线，但不能是纯数字，不能以下划线开头，且不能超过128字符
	VarName string `json:"var_name"`

	Body *UpdateGlobalVariableRequestBody `json:"body,omitempty"`
}

func (o UpdateGlobalVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalVariableRequest struct{}"
	}

	return strings.Join([]string{"UpdateGlobalVariableRequest", string(data)}, " ")
}
