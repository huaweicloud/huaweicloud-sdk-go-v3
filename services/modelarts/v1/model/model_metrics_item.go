package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricsItem 资源指标数据模型。
type MetricsItem struct {
	Table *MetricTableItem `json:"table,omitempty"`

	Metadata *ResourceMetricsMetadata `json:"metadata,omitempty"`
}

func (o MetricsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsItem struct{}"
	}

	return strings.Join([]string{"MetricsItem", string(data)}, " ")
}
