package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLogStreamParams struct {

	// 日志存储时间（天）。
	TtlInDays int32 `json:"ttl_in_days"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`

	// 日志是否存储
	WhetherLogStorage *bool `json:"whether_log_storage,omitempty"`

	// **参数解释：** 是否开启智能冷存储。 **约束限制：** 开启智能冷存储需要日志存储时间大于7天。开启智能冷存储后，需要设置标准存储层数据的保存时间。标准存储至少需保存7天才能转换为智能冷存储层数据。当日志保存时间大于标准存储层数据保存时间，且小于日志存储时间时，日志数据将自动转换为智能冷存储层数据存储。 **取值范围：** - true：开启智能冷存储。 - false：关闭智能冷存储。 **默认取值：** false
	HotColdSeparation *bool `json:"hot_cold_separation,omitempty"`

	// **参数解释：** 日志流标准存储时间（单位：天），即日志上报到LTS后日志标准存储的时间。 **约束限制：** 开启智能冷存储需要日志总存储时间大于7天。 **取值范围：** 最小值：7天 最大值：365天 **默认取值：** 不涉及。
	HotStorageDays *int32 `json:"hot_storage_days,omitempty"`
}

func (o UpdateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogStreamParams struct{}"
	}

	return strings.Join([]string{"UpdateLogStreamParams", string(data)}, " ")
}
