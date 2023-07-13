package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorItemViewConfigResponse Response Object
type ShowMonitorItemViewConfigResponse struct {

	// 标题。
	Title *string `json:"title,omitempty"`

	// 采集器名称。
	CollectorName *string `json:"collector_name,omitempty"`

	// 视图的列表，内部每个List代表的是一行图表。
	ViewRowList *[]ViewRow `json:"view_row_list,omitempty"`

	// 类型。
	Style          *string `json:"style,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMonitorItemViewConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorItemViewConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitorItemViewConfigResponse", string(data)}, " ")
}
