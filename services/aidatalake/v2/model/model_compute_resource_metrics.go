package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComputeResourceMetrics **参数解释**：资源池指标信息。
type ComputeResourceMetrics struct {

	// **参数解释**：数据时间（秒）。 **取值范围**：0~2147483647。
	CollectionTime *int32 `json:"collection_time,omitempty"`

	// **参数解释**：资源池指标值列表，列表元素为资源池指标信息对象。
	Values *[]ComputeMetric `json:"values,omitempty"`
}

func (o ComputeResourceMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeResourceMetrics struct{}"
	}

	return strings.Join([]string{"ComputeResourceMetrics", string(data)}, " ")
}
