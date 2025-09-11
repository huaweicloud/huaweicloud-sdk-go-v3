package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricNameResult struct {

	// **参数解释**: 指标ID。 **取值范围**: 不涉及。
	Metric *string `json:"metric,omitempty"`

	// **参数解释**: 指标名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o MetricNameResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricNameResult struct{}"
	}

	return strings.Join([]string{"MetricNameResult", string(data)}, " ")
}
