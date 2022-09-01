package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备信息
type QueryDeviceSimplifyDto struct {

	// 设备id
	DeviceId *string `json:"device_id,omitempty" xml:"device_id"`

	// 设备识别码
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 父设备id
	GatewayId *string `json:"gateway_id,omitempty" xml:"gateway_id"`

	// 设备名称
	DeviceName *string `json:"device_name,omitempty" xml:"device_name"`

	// 设备协议类型
	ProtocolType *string `json:"protocol_type,omitempty" xml:"protocol_type"`

	// 产品名称
	ProductName *string `json:"product_name,omitempty" xml:"product_name"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`
}

func (o QueryDeviceSimplifyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDeviceSimplifyDto struct{}"
	}

	return strings.Join([]string{"QueryDeviceSimplifyDto", string(data)}, " ")
}
