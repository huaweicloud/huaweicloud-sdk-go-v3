package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 全局参数
type Parameter struct {

	// 参数名称,需要满足：[^\\>+<;#\";&?%=']{1,64}。
	ParamName string `json:"param_name"`

	// 参数类型。
	ParamType *string `json:"param_type,omitempty"`

	// 参数分组。
	ParamGroup *string `json:"param_group,omitempty"`

	// 参数初始值。
	DefaultValue *string `json:"default_value,omitempty"`

	// 参数id。
	Id *string `json:"id,omitempty"`

	// 是否加密。
	Encryption bool `json:"encryption"`

	// 参数提示。
	Hint *string `json:"hint,omitempty"`

	// 是否从参数库选择。
	QuoteParam bool `json:"quote_param"`

	// 是否为必填参数。
	Required bool `json:"required"`

	// 参数描述。
	Description *string `json:"description,omitempty"`
}

func (o Parameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Parameter struct{}"
	}

	return strings.Join([]string{"Parameter", string(data)}, " ")
}
