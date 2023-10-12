package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDimensionName 维度名，多个维度按字母序逗号分开
type MetricDimensionName struct {
}

func (o MetricDimensionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimensionName struct{}"
	}

	return strings.Join([]string{"MetricDimensionName", string(data)}, " ")
}
