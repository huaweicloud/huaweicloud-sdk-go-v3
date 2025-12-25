package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDimensionSchema **参数解释**: 资源维度信息, 多维度逗号分隔。 **约束限制**: 不涉及。 **取值范围**: 必须以字母开头，只能包含0-9/a-z/A-Z/_/-/,。每个维度必须以字母开头，每个维度长度最短1，最长32，多个维度直接用,分隔。 **默认取值**: 不涉及。
type MetricDimensionSchema struct {
}

func (o MetricDimensionSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimensionSchema struct{}"
	}

	return strings.Join([]string{"MetricDimensionSchema", string(data)}, " ")
}
