package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDevicesResponse struct {

	// 查询设备列表响应结构体
	Devices *[]QueryDeviceSimplifyDto `json:"devices,omitempty" xml:"devices"`

	// 满足查询条件的记录总数。
	Count *int64 `json:"count,omitempty" xml:"count"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListDevicesResponse", string(data)}, " ")
}
