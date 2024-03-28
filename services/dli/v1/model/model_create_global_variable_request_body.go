package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalVariableRequestBody 创建全局变量
type CreateGlobalVariableRequestBody struct {

	// 变量名称
	VarName string `json:"var_name"`

	// 变量的值
	VarValue string `json:"var_value"`

	// 是否为敏感变量
	IsSensitive *bool `json:"is_sensitive,omitempty"`
}

func (o CreateGlobalVariableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalVariableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGlobalVariableRequestBody", string(data)}, " ")
}
