package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListLogStreamsResponseBody1LogStreams struct {

	// 日志流创建时间
	CreationTime int64 `json:"creation_time"`

	// 日志流ID
	LogStreamId string `json:"log_stream_id"`

	// 日志流名称
	LogStreamName string `json:"log_stream_name"`

	// 日志流别名
	LogStreamNameAlias *string `json:"log_stream_name_alias,omitempty"`

	// 日志流所属标签
	Tag map[string]string `json:"tag"`

	// 过滤器个数
	FilterCount int32 `json:"filter_count"`

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

	// 日志组ID
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释：** 是否收藏日志流。 **取值范围：** - true：收藏日志流。 - false：不收藏日志流。
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

func (o ListLogStreamsResponseBody1LogStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamsResponseBody1LogStreams struct{}"
	}

	return strings.Join([]string{"ListLogStreamsResponseBody1LogStreams", string(data)}, " ")
}
