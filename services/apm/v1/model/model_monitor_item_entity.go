package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonitorItemEntity struct {

	// 采集器类别id。
	CategoryId *int32 `json:"category_id,omitempty"`

	// 采集器名称。
	CollectorName *string `json:"collector_name,omitempty"`

	// 采集器类别展示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 是否展示标题。
	ShowInTotal *bool `json:"show_in_total,omitempty"`

	// 监控项id。
	MonitorItemId *int64 `json:"monitor_item_id,omitempty"`

	// 是否禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// 采集器id。
	CollectorId *int32 `json:"collector_id,omitempty"`

	// 序列号。
	Sequence *int32 `json:"sequence,omitempty"`

	// 默认数据采集间隔。
	CollectInterval *int32 `json:"collect_interval,omitempty"`
}

func (o MonitorItemEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorItemEntity struct{}"
	}

	return strings.Join([]string{"MonitorItemEntity", string(data)}, " ")
}
