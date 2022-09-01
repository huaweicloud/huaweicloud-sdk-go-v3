package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEnvMonitorItemsResponse struct {

	// 采集器类别集合
	CategoryInfoList *[]CollectorCategoryInfo `json:"category_info_list,omitempty" xml:"category_info_list"`

	// 监控项集合
	MonitorItemInfoList *[]MonitorItemEntity `json:"monitor_item_info_list,omitempty" xml:"monitor_item_info_list"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowEnvMonitorItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvMonitorItemsResponse struct{}"
	}

	return strings.Join([]string{"ShowEnvMonitorItemsResponse", string(data)}, " ")
}
