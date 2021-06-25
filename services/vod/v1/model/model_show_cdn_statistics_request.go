package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCdnStatisticsRequest struct {
	// 开始时间

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间

	EndTime *string `json:"end_time,omitempty"`
	// 支持的参数类型 cdn_bw：CDN峰值带宽cdn_flux：CDN流量req_num：请求总数req_hit_rate：请求命中率flux_hit_rate：流量命中率

	StatType string `json:"stat_type"`
	// 域名列表，多个域名以逗号（半角）分隔

	Domain string `json:"domain"`
	// 采样间隔，单位：秒，取值说明： 时间跨度1天：5分钟、1小时、4小时、8小时，分别对应300秒、3600秒、14400秒和28800秒。 时间跨度2~7天：1小时、4小时、8小时、1天，分别对应3600秒、14400秒、28800秒和86400秒。 时间跨度8~31天：4小时、8小时、1天，分别对应14400秒、28800秒和86400秒。 如果不传，默认取对应时间跨度的最小间隔。

	Interval *int32 `json:"interval,omitempty"`
}

func (o ShowCdnStatisticsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCdnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowCdnStatisticsRequest", string(data)}, " ")
}
