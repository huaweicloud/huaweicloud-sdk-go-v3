package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvMonitorItemsResponse Response Object
type ShowEnvMonitorItemsResponse struct {

	// 采集器类别集合。
	CategoryInfoList *[]CollectorCategoryInfo `json:"category_info_list,omitempty"`

	// 监控项集合。
	MonitorItemInfoList *[]MonitorItemEntity `json:"monitor_item_info_list,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowEnvMonitorItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvMonitorItemsResponse struct{}"
	}

	return strings.Join([]string{"ShowEnvMonitorItemsResponse", string(data)}, " ")
}
