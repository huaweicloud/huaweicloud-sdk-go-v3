package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerMetricDataRequest Request Object
type ShowServerMetricDataRequest struct {

	// 服务器唯一标识。
	ServerId string `json:"server_id"`

	// 监控开始时间：由日期加时间组成，UTC格式，例如“2021-05-11T11:45:42.000Z”。
	StartTime *sdktime.SdkTime `json:"start_time"`

	// 监控结束时间：由日期加时间组成，UTC格式，例如“2021-05-11T11:45:42.000Z”。
	EndTime *sdktime.SdkTime `json:"end_time"`

	// 监控查询指标：例如 \"cpu_util\"，详情见下表；当metric_name为空时，为批量查询。| 指标                                  | 指标名称          | 测量对象     | 监控周期 || ------------------------------------- | ----------------- | ------------ | -------- || cpu_util                              | CPU使用率         | 弹性云服务器 | 5分钟    || mem_util                              | 内存使用率        | 弹性云服务器 | 5分钟    || disk_util_inband                      | 磁盘使用率        | 弹性云服务器 | 5分钟    || disk_read_bytes_rate                  | 磁盘读带宽        | 弹性云服务器 | 5分钟    || disk_write_bytes_rate                 | 磁盘写带宽        | 弹性云服务器 | 5分钟    || disk_read_requests_rate               | 磁盘读IOPS        | 弹性云服务器 | 5分钟    || disk_write_requests_rate              | 磁盘写IOPS        | 弹性云服务器 | 5分钟    || network_incoming_bytes_rate_inband    | 带内网络流入速率  | 弹性云服务器 | 5分钟    || network_outgoing_bytes_rate_inband    | 带内网络流出速率  | 弹性云服务器 | 5分钟    || network_incoming_bytes_aggregate_rate | 带外网络流入速率  | 弹性云服务器 | 5分钟    || network_outgoing_bytes_aggregate_rate | 带外网络流出速率  | 弹性云服务器 | 5分钟    |
	MetricName *string `json:"metric_name,omitempty"`
}

func (o ShowServerMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ShowServerMetricDataRequest", string(data)}, " ")
}
