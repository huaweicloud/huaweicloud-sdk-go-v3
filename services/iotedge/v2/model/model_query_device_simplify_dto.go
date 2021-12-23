package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备信息
type QueryDeviceSimplifyDto struct {
	// 设备id

	DeviceId *string `json:"device_id,omitempty"`
	// 设备识别码

	NodeId *string `json:"node_id,omitempty"`
	// 父设备id

	GatewayId *string `json:"gateway_id,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 设备协议类型

	ProtocolType *string `json:"protocol_type,omitempty"`
	// 产品名称

	ProductName *string `json:"product_name,omitempty"`
	// 产品ID

	ProductId *string `json:"product_id,omitempty"`
}

func (o QueryDeviceSimplifyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDeviceSimplifyDto struct{}"
	}

	return strings.Join([]string{"QueryDeviceSimplifyDto", string(data)}, " ")
}
