package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDevicesResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回数量

	Size *int32 `json:"size,omitempty"`
	// 设备接入地址

	ConnectAddress *string `json:"connect_address,omitempty"`
	// 设备接入SSL地址

	SslConnectAddress *string `json:"ssl_connect_address,omitempty"`
	// 设备接入IPV6地址

	Ipv6ConnectAddress *string `json:"ipv6_connect_address,omitempty"`
	// 设备接入IPV6 SSL地址

	Ipv6SslConnectAddress *string `json:"ipv6_ssl_connect_address,omitempty"`
	// 设备ID列表

	Items          *[]Device `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListDevicesResponse", string(data)}, " ")
}
