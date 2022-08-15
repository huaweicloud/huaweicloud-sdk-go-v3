package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonitorItemEntity struct {
	CategoryId *int32 `json:"category_id,omitempty"`

	CollectorName *string `json:"collector_name,omitempty"`

	DisplayName *string `json:"display_name,omitempty"`

	ShowInTotal *bool `json:"show_in_total,omitempty"`

	MonitorItemId *int64 `json:"monitor_item_id,omitempty"`

	Disabled *bool `json:"disabled,omitempty"`

	CollectorId *int32 `json:"collector_id,omitempty"`

	Sequence *int32 `json:"sequence,omitempty"`

	CollectInterval *int32 `json:"collect_interval,omitempty"`
}

func (o MonitorItemEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorItemEntity struct{}"
	}

	return strings.Join([]string{"MonitorItemEntity", string(data)}, " ")
}
