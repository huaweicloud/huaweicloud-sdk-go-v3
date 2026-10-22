package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OverviewStatistics 总览统计数据
type OverviewStatistics struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 用户输入请求数
	TextRequestCount int64 `json:"text_request_count"`

	// 用户输出请求数
	ImageRequestCount int64 `json:"image_request_count"`

	// 按接入类型统计
	AccessTypeStats []AccessTypeStatistics `json:"access_type_stats"`

	// 分类别统计数据
	CategoryStats []CategoryStatistics `json:"category_stats"`

	// 处置动作统计数据
	ActionStats []ActionStatistics `json:"action_stats"`

	TimeRange *TimeRange `json:"time_range"`

	// 已开启的AI应用数量
	ActiveAppCount *int64 `json:"active_app_count,omitempty"`
}

func (o OverviewStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverviewStatistics struct{}"
	}

	return strings.Join([]string{"OverviewStatistics", string(data)}, " ")
}
