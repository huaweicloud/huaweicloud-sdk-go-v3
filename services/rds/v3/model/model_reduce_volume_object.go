package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReduceVolumeObject 实例磁盘缩容时必填。
type ReduceVolumeObject struct {

	// 缩容后实例磁盘的目标大小。每次缩容至少缩小10GB；目标大小必须为10的整数倍。 为确保实例的正常使用，根据当前磁盘的使用量情况存在磁盘容量下限，当此参数小于磁盘容量下限时，缩容会下发失败，此时请适当调大此参数。
	Size int32 `json:"size"`

	// 是否定时变更。 - true，为定时在运维时间窗做变更。 - false，为立即变更，默认该方式。
	IsDelay bool `json:"is_delay"`

	// 该参数只有磁盘类型为Flexible SSD（GPSSD2）和极速型SSDV2（ESSD2）的磁盘必填。 对于Flexible SSD类型的磁盘，IOPS值配置的范围为3000~128000，具体可配置值受磁盘大小限制，需要小于等于500*磁盘容量。 对于极速型SSDV2类型的磁盘，IOPS值配置的范围为100~256000，具体可配置值受磁盘大小限制，需要小于等于1000*磁盘容量。
	Iops *int32 `json:"iops,omitempty"`

	// 该参数只有磁盘类型为Flexible SSD（GPSSD2）的磁盘必填。 对于Flexible SSD类型的磁盘，throughput值配置的范围为125~1000 MiB/s，具体可配置值受IOPS大小限制，需要小于等于IOPS/4。
	Throughput *int32 `json:"throughput,omitempty"`
}

func (o ReduceVolumeObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReduceVolumeObject struct{}"
	}

	return strings.Join([]string{"ReduceVolumeObject", string(data)}, " ")
}
