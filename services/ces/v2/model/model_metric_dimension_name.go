package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDimensionName **参数解释** 维度名，多个维度按字母序逗号分开 **约束限制** 不涉及 **取值范围** 字符串长度[1,131] **默认取值** 不涉及
type MetricDimensionName struct {
}

func (o MetricDimensionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimensionName struct{}"
	}

	return strings.Join([]string{"MetricDimensionName", string(data)}, " ")
}
