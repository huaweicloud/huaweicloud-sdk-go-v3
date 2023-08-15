package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DifferentDetails struct {

	// 参数名称
	ParameterName string `json:"parameter_name"`

	// 源参数模板中的参数值。
	SourceValue string `json:"source_value"`

	// 目标参数模板中的参数值。
	TargetValue string `json:"target_value"`
}

func (o DifferentDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DifferentDetails struct{}"
	}

	return strings.Join([]string{"DifferentDetails", string(data)}, " ")
}
