package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBuildJobParameterParam 构建执行参数子参数
type UpdateBuildJobParameterParam struct {

	// 参数字段名
	Name *string `json:"name,omitempty"`

	// 参数字段值
	Value *string `json:"value,omitempty"`

	// 枚举类参数限制
	Limits *[]LimitsParam `json:"limits,omitempty"`
}

func (o UpdateBuildJobParameterParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBuildJobParameterParam struct{}"
	}

	return strings.Join([]string{"UpdateBuildJobParameterParam", string(data)}, " ")
}
