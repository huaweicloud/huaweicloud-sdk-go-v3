package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnStatisticsRequest Request Object
type ListCdnStatisticsRequest struct {

	// 开始时间，格式为yyyymmddhhmmss。   - interval为300时，end_time设置为整5分钟时刻点，如：20240601000000。   - interval为3600时，end_time设置为整小时时刻点，如：20240601120000。   - interval为86400时，end_time设置为东8区零点时刻点，如：20240601000000。   - 只能查询最近三个月内的数据，且时间跨度不能超过31天。
	StartTime string `json:"start_time"`

	// 结束时间，格式为yyyymmddhhmmss。   - interval为300时，end_time设置为整5分钟时刻点，如：20240601000000。   - interval为3600时，end_time设置为整小时时刻点，如：20240601120000。   - interval为86400时，end_time设置为东8区零点时刻点，如：20240601000000。   - 只能查询最近三个月内的数据，且时间跨度不能超过31天。
	EndTime string `json:"end_time"`

	// 统计数据类型。  取值如下： - bw：CDN峰值带宽 - flux：CDN流量 - req_num：请求总数 - req_hit_rate：请求命中率 - flux_hit_rate：流量命中率 - http_code_2xx 状态码汇总2xx - http_code_3xx 状态码汇总3xx - http_code_4xx 状态码汇总4xx - http_code_5xx 状态码汇总5xx  每次只能查询一种统计数据。
	StatType string `json:"stat_type"`

	// 域名列表，多个域名以逗号（半角）分隔。  示例：example.test1.com,example.test2.com。  ALL表示查询名下全部域名。一次最多查询20个域名。
	Domain string `json:"domain"`

	// 查询粒度间隔。  取值如下： - 300(5分钟)：最大查询跨度2天。 - 3600(1小时)：最大查询跨度7天。 - 86400(1天)：最大查询跨度31天，最少跨度为2天。  单位：秒。  若不设置，小于1天 300，大于1天小于7天3600，大于7天86400。
	Interval *int32 `json:"interval,omitempty"`
}

func (o ListCdnStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListCdnStatisticsRequest", string(data)}, " ")
}
