package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnvMonitorItemResponse struct {

	// 监控项列表
	MonitorItemList *[]MonitorItem `json:"monitor_item_list,omitempty" xml:"monitor_item_list"`

	// 总数
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 总页数
	TotalPage      *int32 `json:"totalPage,omitempty" xml:"totalPage"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnvMonitorItemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvMonitorItemResponse struct{}"
	}

	return strings.Join([]string{"ListEnvMonitorItemResponse", string(data)}, " ")
}
