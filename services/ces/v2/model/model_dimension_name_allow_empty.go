package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimensionNameAllowEmpty 资源维度，必须以字母开头，多维度用\",\"分割，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32, 事件告警模板DimensionName为空
type DimensionNameAllowEmpty struct {
}

func (o DimensionNameAllowEmpty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionNameAllowEmpty struct{}"
	}

	return strings.Join([]string{"DimensionNameAllowEmpty", string(data)}, " ")
}
