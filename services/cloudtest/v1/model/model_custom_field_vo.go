package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomFieldVo 自定义字段信息
type CustomFieldVo struct {

	// 字段类型（单行文本text、多行文本textArea、单选框radio、多选框checkBox、日期date、数字number、单选用户user）
	Type *string `json:"type,omitempty"`

	// 测试用例自定义字段值
	Value *string `json:"value,omitempty"`

	// 项目用例自定义字段入参或者返回参数名称
	CustomFieldParam *string `json:"custom_field_param,omitempty"`

	// user类型测试用例自定义字段对应用户名，其它类型字段不返回
	UserName *string `json:"user_name,omitempty"`
}

func (o CustomFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomFieldVo struct{}"
	}

	return strings.Join([]string{"CustomFieldVo", string(data)}, " ")
}
