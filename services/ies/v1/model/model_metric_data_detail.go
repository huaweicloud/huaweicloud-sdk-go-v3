package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricDataDetail struct {

	// 监控指标名称，当前包含指标： - cpu_used：该维度vCPU已使用核数，单位：个，支持维度：site_id，flavor - cpu_available_total：用户可使用该维度vCPU总核数，单位：个，支持维度：site_id，flavor - cpu_total：该维度vCPU总核数（包含HA等预留核数），单位：个，支持维度：site_id，flavor - memory_used：该维度内存已使用量，单位：Gb，支持维度：site_id，flavor - memory_available_total：用户可使用该维度内存总量，单位：Gb，支持维度：site_id，flavor - memory_total：该维度内存总量（包含HA等预留内存量），单位：Gb，支持维度：site_id，flavor - capacity_used：该维度块存储资源已使用量，单位：GiB，支持维度：site_id，storage - capacity_available_total：用户可使用该维度块存储资源总容量（用户订购开通的存储容量），单位：GiB，支持维度：site_id，storage - capacity_total：当前已订购的资源场景下该维度块存储资源最大容量（订购资源包含的存储容量可能大于用户已开通容量），单位：GiB，支持维度：site_id，storage - available：该维度对应规格剩余可发放数量，单位：台，支持维度：flavor_capacity
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
