package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogStreamResBody 返回的日志流信息
type LogStreamResBody struct {

	// 创建时间 最小值：1577808000000 最大值：4102416000000
	CreationTime *int64 `json:"creation_time,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 日志流别名
	LogStreamNameAlias *string `json:"log_stream_name_alias,omitempty"`

	// 日志流所属标签
	Tag map[string]string `json:"tag,omitempty"`

	// 过滤器个数
	FilterCount *int32 `json:"filter_count,omitempty"`

	// 是否收藏日志流。
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// 是否日志存储
	WhetherLogStorage *bool `json:"whether_log_storage,omitempty"`

	// 是否冷存储
	HotColdSeparation *bool `json:"hot_cold_separation,omitempty"`

	// 匿名写入开关
	AuthWebTracking *bool `json:"auth_web_tracking,omitempty"`

	// 存储时间
	TtlInDays *int32 `json:"ttl_in_days,omitempty"`

	// 标准存储时间
	HotStorageDays *int32 `json:"hot_storage_days,omitempty"`
}

func (o LogStreamResBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogStreamResBody struct{}"
	}

	return strings.Join([]string{"LogStreamResBody", string(data)}, " ")
}
