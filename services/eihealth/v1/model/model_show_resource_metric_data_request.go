package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResourceMetricDataRequest struct {

	// 查询监控数据起始时间，UNIX时间戳，单位毫秒，不填时默认为当前时间
	FromTime *int64 `json:"from_time,omitempty"`

	// 查询数据截止时间，UNIX时间戳，单位毫秒，不填时默认为当前时间
	ToTime *int64 `json:"to_time,omitempty"`

	// 监控数据周期。枚举值，取值范围：real_time（实时数据）、five_minutes（5分钟粒度）、fifteen_to_twenty_minutes（15-20分钟粒度）、one_hour（1小时粒度），不填时默认为real_time
	Period *string `json:"period,omitempty"`

	// 统计方法。枚举值，取值范围：max（最大值）、min（最小值）、sum（求和值）、average（平均值），不填时默认为max
	Method *string `json:"method,omitempty"`

	// 查询的监控指标名称
	MetricName string `json:"metric_name"`

	// 查询的监控资源对象id，当查询存储资源和计算节点资源中的集群监控数据时，不需要填写资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 显卡id，仅查询裸金属节点的gpu监控时，需要指定
	DeviceId *string `json:"device_id,omitempty"`
}

func (o ShowResourceMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceMetricDataRequest", string(data)}, " ")
}
