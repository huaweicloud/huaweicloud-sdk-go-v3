package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParameterValuesInfo struct {

	// 是否需要重启
	RestartRequired *bool `json:"restart_required,omitempty"`

	// 是否只读
	Readonly *bool `json:"readonly,omitempty"`

	// 参数名
	Name *string `json:"name,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`

	// 参数类型
	Type *string `json:"type,omitempty"`

	// 参数描述
	Description *string `json:"description,omitempty"`
}

func (o ParameterValuesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterValuesInfo struct{}"
	}

	return strings.Join([]string{"ParameterValuesInfo", string(data)}, " ")
}
