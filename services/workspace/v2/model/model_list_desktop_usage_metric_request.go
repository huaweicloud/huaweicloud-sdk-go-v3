package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopUsageMetricRequest Request Object
type ListDesktopUsageMetricRequest struct {

	// 查询起始时间(0时区) 云服务每天凌晨02:00进行聚合运算前一天00:00:00~23:59:59的使用时长,并将周期范围内的数据聚合到周期边界上 跨天的记录会按照统计周期进行计算 假设一天内桌面登录多次，09:00~12:00,13:00~21:00,22:00~01:00(次日): 则当天的累计使用时长数据会被汇聚到23:59:59这个点;总使用时长为 3hours(09:00~12:00)+8hours(13:00~21:00)+2hours(22:00~00:00) 如果查询的from-to不足一个周期内，可能造成查询到数据为空；
	StartTime string `json:"start_time"`

	// 查询截至时间(0时区)
	EndTime string `json:"end_time"`

	// 资源名称(模糊匹配)
	ResourceName *string `json:"resource_name,omitempty"`

	// 最小空闲天数
	MinIdleDays *int32 `json:"min_idle_days,omitempty"`

	// 最大空闲天数 min_idle_days、max_idle_days都非空时,max_idle_days必须大于等于min_idle_days否则可能查询不到数据
	MaxIdleDays *int32 `json:"max_idle_days,omitempty"`

	// 使用时长(hour)最小值
	UsageMinHours *int32 `json:"usage_min_hours,omitempty"`

	// 使用时长(hour)最大值(必须大于等于usage_min_hours)
	UsageMaxHours *int32 `json:"usage_max_hours,omitempty"`

	// 按照指标进行排序 * `desktop_usage` -  按照桌面使用时长排序 * `desktop_idle_duration` -  按照桌面空闲周期排序
	SortField *string `json:"sort_field,omitempty"`

	// 按照指标进行排序的方向;需配合sort_field起义使用 * `DESC` - 降序返回数据 * `ASC` -  升序返回数据
	SortType *string `json:"sort_type,omitempty"`

	// 查询的偏移量,默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100],默认值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopUsageMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopUsageMetricRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopUsageMetricRequest", string(data)}, " ")
}
