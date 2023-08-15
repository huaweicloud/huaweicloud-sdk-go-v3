package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffDetails struct {

	// 参数名称
	ParameterName string `json:"parameter_name"`

	// 比较参数模板的参数值。
	SourceValue string `json:"source_value"`

	// 目标参数模板的参数值。
	TargetValue string `json:"target_value"`
}

func (o DiffDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffDetails struct{}"
	}

	return strings.Join([]string{"DiffDetails", string(data)}, " ")
}
