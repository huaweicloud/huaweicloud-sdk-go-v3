package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParamGroupParameterDifferences struct {

	// 参数名称。
	ParameterName *string `json:"parameter_name,omitempty"`

	// 源参数模板中的参数值。
	SourceValue *string `json:"source_value,omitempty"`

	// 目标参数模板中的参数值。
	TargetValue *string `json:"target_value,omitempty"`
}

func (o ParamGroupParameterDifferences) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupParameterDifferences struct{}"
	}

	return strings.Join([]string{"ParamGroupParameterDifferences", string(data)}, " ")
}
