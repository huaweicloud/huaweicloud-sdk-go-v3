package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevicesResponse Response Object
type ListDevicesResponse struct {

	// 查询设备列表响应结构体
	Devices *[]QueryDeviceSimplifyDto `json:"devices,omitempty"`

	// 满足查询条件的记录总数。
	Count *int64 `json:"count,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListDevicesResponse", string(data)}, " ")
}
