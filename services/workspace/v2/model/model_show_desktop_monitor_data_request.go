package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopMonitorDataRequest Request Object
type ShowDesktopMonitorDataRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	// 监控开始时间：由日期加时间组成，UTC格式，例如“2021-05-11T11:45:42Z”。
	StartTime string `json:"start_time"`

	// 监控结束时间：由日期加时间组成，UTC格式，例如“2021-05-11T11:45:42Z”。
	EndTime string `json:"end_time"`

	// 监控查询指标：例如 \"cpu_util\"，详情见下表；当metric_name为空时，为批量查询| 指标                                  | 指标名称          | 测量对象     | 监控周期 || ------------------------------------- | ----------------- | ------------ | -------- || cpu_util                              | CPU使用率         |    云桌面    | 5分钟    || mem_util                              | 内存使用率        |    云桌面    | 5分钟    || disk_util_inband                      | 磁盘使用率        |    云桌面    | 5分钟    || disk_read_bytes_rate                  | 磁盘读带宽        |    云桌面    | 5分钟    || disk_write_bytes_rate                 | 磁盘写带宽        |    云桌面    | 5分钟    || disk_read_requests_rate               | 磁盘读IOPS        |    云桌面    | 5分钟    || disk_write_requests_rate              | 磁盘写IOPS        |    云桌面    | 5分钟    || network_incoming_bytes_rate_inband    | 带内网络流入速率  |    云桌面    | 5分钟    || network_outgoing_bytes_rate_inband    | 带内网络流出速率  |    云桌面    | 5分钟    || network_incoming_bytes_aggregate_rate | 带外网络流入速率  |    云桌面    | 5分钟    || network_outgoing_bytes_aggregate_rate | 带外网络流出速率  |    云桌面    | 5分钟    || network_vm_connections                | 网络连接数        |    云桌面    | 5分钟    || cpu_credit_usage                      | CPU积分使用量     |    云桌面    | 5分钟    || cpu_credit_balance                    | CPU积分累积量     |    云桌面    | 5分钟    || cpu_surplus_credit_balance            | CPU超额积分累积量 |    云桌面    | 5分钟    || cpu_surplus_credit_charged            | CPU超额积分收费量 |    云桌面    | 5分钟    || user_online_info                      | 用户在线信息      |    云桌面    | -        |
	MetricName *string `json:"metric_name,omitempty"`
}

func (o ShowDesktopMonitorDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopMonitorDataRequest struct{}"
	}

	return strings.Join([]string{"ShowDesktopMonitorDataRequest", string(data)}, " ")
}
