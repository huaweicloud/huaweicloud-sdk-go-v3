package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDimensionStrResp **参数解释**: 指标维度, 多维度逗号分隔。。 **取值范围**: 必须以字母开头，只能包含0-9/a-z/A-Z/_/-/,。每个维度必须以字母开头，每个维度长度最短1，最长32，多个维度直接用,分割。
type MetricDimensionStrResp struct {
}

func (o MetricDimensionStrResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDimensionStrResp struct{}"
	}

	return strings.Join([]string{"MetricDimensionStrResp", string(data)}, " ")
}
