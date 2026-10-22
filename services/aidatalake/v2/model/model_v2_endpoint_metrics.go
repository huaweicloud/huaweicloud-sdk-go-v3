package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2EndpointMetrics **参数解释**：端点指标信息。
type V2EndpointMetrics struct {

	// **参数解释**：数据时间（秒）。 **取值范围**：不涉及。
	CollectionTime *int32 `json:"collection_time,omitempty"`

	// **参数解释**：端点指标值列表，列表元素为端点指标信息对象。
	Values *[]V2EndpointMetric `json:"values,omitempty"`
}

func (o V2EndpointMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointMetrics struct{}"
	}

	return strings.Join([]string{"V2EndpointMetrics", string(data)}, " ")
}
