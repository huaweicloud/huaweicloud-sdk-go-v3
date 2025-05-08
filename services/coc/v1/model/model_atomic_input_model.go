package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AtomicInputModel struct {

	// 参数变量名，执行时原子能力内引用
	ParamKey *string `json:"param_key,omitempty"`

	// 参数中文名，下拉展示时使用
	ParamNameZh *string `json:"param_name_zh,omitempty"`

	// 参数英文名，下拉展示时使用
	ParamNameEn *string `json:"param_name_en,omitempty"`

	// 是否必填
	Required *bool `json:"required,omitempty"`

	// 参数类型：常量/字典
	ParamType *string `json:"param_type,omitempty"`

	// 最小值
	Min *int32 `json:"min,omitempty"`

	// 最大值
	Max *int32 `json:"max,omitempty"`

	// 长度最小值
	MinLen *int32 `json:"min_len,omitempty"`

	// 长度最大值
	MaxLen *int32 `json:"max_len,omitempty"`

	// 正则限制表达式
	Pattern *string `json:"pattern,omitempty"`
}

func (o AtomicInputModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtomicInputModel struct{}"
	}

	return strings.Join([]string{"AtomicInputModel", string(data)}, " ")
}
