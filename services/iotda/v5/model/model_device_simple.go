package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceSimple 设备简单信息结构体
type DeviceSimple struct {

	// 资源空间ID。
	AppId *string `json:"app_id,omitempty"`

	// 设备ID，用于唯一标识一个设备。在注册设备时直接指定，或者由物联网平台分配获得。由物联网平台分配时，生成规则为\"product_id\" + \"_\" + \"node_id\"拼接而成。
	DeviceId *string `json:"device_id,omitempty"`

	// 设备标识码，通常使用IMEI、MAC地址或Serial No作为node_id。
	NodeId *string `json:"node_id,omitempty"`

	// 网关ID，用于标识设备所属的父设备，即父设备的设备ID。当设备是直连设备时，gateway_id与设备的device_id一致。当设备是非直连设备时，gateway_id为设备所关联的父设备的device_id。
	GatewayId *string `json:"gateway_id,omitempty"`

	// 设备名称。
	DeviceName *string `json:"device_name,omitempty"`

	// 设备关联的产品ID，用于唯一标识一个产品模型。
	ProductId *string `json:"product_id,omitempty"`
}

func (o DeviceSimple) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceSimple struct{}"
	}

	return strings.Join([]string{"DeviceSimple", string(data)}, " ")
}
