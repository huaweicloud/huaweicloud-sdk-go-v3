package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiTemplateVariable struct {

	// 变量id
	Id *int64 `json:"id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 模板名称
	TempName *string `json:"temp_name,omitempty"`

	// 变量索引
	VariableIndex *int32 `json:"variable_index,omitempty"`

	// 变量类型
	VariableType *string `json:"variable_type,omitempty"`

	// 变量长度
	VariableLength *int32 `json:"variable_length,omitempty"`

	// 变量描述
	VariableDesc *string `json:"variable_desc,omitempty"`
}

func (o ApiTemplateVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiTemplateVariable struct{}"
	}

	return strings.Join([]string{"ApiTemplateVariable", string(data)}, " ")
}
