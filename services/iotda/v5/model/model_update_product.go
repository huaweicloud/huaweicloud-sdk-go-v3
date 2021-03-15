package model

import (
	"encoding/json"

	"strings"
)

// 修改产品信息结构体。
type UpdateProduct struct {
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，必须携带该参数指定修改的产品属于哪个资源空间，否则接口会提示错误。如果用户存在多资源空间，同时又不想携带该参数，可以联系华为技术支持对用户数据做资源空间合并。
	AppId *string `json:"app_id,omitempty"`
	// 产品名称。
	Name *string `json:"name,omitempty"`
	// 设备类型。
	DeviceType *string `json:"device_type,omitempty"`
	// 设备使用的协议类型。取值范围：MQTT，CoAP，HTTP，HTTPS，Modbus，ONVIF， OPC-UA。
	ProtocolType *string `json:"protocol_type,omitempty"`
	// 设备上报数据的格式，取值范围：json，binary。
	DataFormat *string `json:"data_format,omitempty"`
	// 设备的服务能力列表。
	ServiceCapabilities *[]ServiceCapability `json:"service_capabilities,omitempty"`
	// 厂商名称。
	ManufacturerName *string `json:"manufacturer_name,omitempty"`
	// 设备所属行业。
	Industry *string `json:"industry,omitempty"`
	// 产品的描述信息。
	Description *string `json:"description,omitempty"`
}

func (o UpdateProduct) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProduct struct{}"
	}

	return strings.Join([]string{"UpdateProduct", string(data)}, " ")
}
