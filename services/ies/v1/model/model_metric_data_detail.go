package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricDataDetail struct {

	// 监控指标名称
	Name *string `json:"name,omitempty"`

	// 监控值
	Value *int64 `json:"value,omitempty"`

	// 记录更新时间
	ReadAt *sdktime.SdkTime `json:"read_at,omitempty"`

	Dimension *MetricDataDetailDimension `json:"dimension,omitempty"`
}

func (o MetricDataDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataDetail struct{}"
	}

	return strings.Join([]string{"MetricDataDetail", string(data)}, " ")
}
