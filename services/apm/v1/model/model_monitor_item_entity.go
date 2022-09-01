package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonitorItemEntity struct {
	CategoryId *int32 `json:"category_id,omitempty" xml:"category_id"`

	CollectorName *string `json:"collector_name,omitempty" xml:"collector_name"`

	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	ShowInTotal *bool `json:"show_in_total,omitempty" xml:"show_in_total"`

	MonitorItemId *int64 `json:"monitor_item_id,omitempty" xml:"monitor_item_id"`

	Disabled *bool `json:"disabled,omitempty" xml:"disabled"`

	CollectorId *int32 `json:"collector_id,omitempty" xml:"collector_id"`

	Sequence *int32 `json:"sequence,omitempty" xml:"sequence"`

	CollectInterval *int32 `json:"collect_interval,omitempty" xml:"collect_interval"`
}

func (o MonitorItemEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorItemEntity struct{}"
	}

	return strings.Join([]string{"MonitorItemEntity", string(data)}, " ")
}
