package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorItemDetailResponse Response Object
type ShowMonitorItemDetailResponse struct {

	// 采集间隔
	Interval *int64 `json:"interval,omitempty"`

	// 采集器ID
	CollectorId *int64 `json:"collector_id,omitempty"`

	// 采集参数配置列表
	ConfigItemList *[]ConfigItemValue `json:"config_item_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowMonitorItemDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorItemDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitorItemDetailResponse", string(data)}, " ")
}
