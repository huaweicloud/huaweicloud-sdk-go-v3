package model

import (
	"encoding/json"

	"strings"
)

// 创建产品信息结构体。
type AddProduct struct {
	// 产品ID，用于唯一标识一个产品。如果携带此参数，平台将产品ID设置为该参数值；如果不携带此参数，产品ID在物联网平台创建产品后由平台分配获得。
	ProductId *string `json:"product_id,omitempty"`
	// 产品名称。
	Name string `json:"name"`
	// 设备类型。
	DeviceType string `json:"device_type"`
	// 设备使用的协议类型。取值范围：MQTT，CoAP，HTTP，HTTPS，Modbus，ONVIF， OPC-UA。
	ProtocolType string `json:"protocol_type"`
	// 设备上报数据的格式，取值范围：json，binary。默认值json。
	DataFormat string `json:"data_format"`
	// 设备的服务能力列表。
	ServiceCapabilities []ServiceCapability `json:"service_capabilities"`
	// 厂商名称。
	ManufacturerName *string `json:"manufacturer_name,omitempty"`
	// 设备所属行业。
	Industry *string `json:"industry,omitempty"`
	// 产品的描述信息。
	Description *string `json:"description,omitempty"`
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的产品归属到哪个资源空间下，否则创建的产品将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。
	AppId *string `json:"app_id,omitempty"`
}

func (o AddProduct) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddProduct struct{}"
	}

	return strings.Join([]string{"AddProduct", string(data)}, " ")
}
