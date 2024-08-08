package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Port 企业版端口信息。创建企业版实例时必填。
type Port struct {

	// **参数说明**：应用接入HTTPS协议端口，默认值：443
	AppHttpsPort *int32 `json:"app_https_port,omitempty"`

	// **参数说明**：应用接入AMQP协议端口, 默认值：5671
	AppAmqpsPort *int32 `json:"app_amqps_port,omitempty"`

	// **参数说明**：设备接入COAP协议端口, 默认值：5683
	DeviceCoapPort *int32 `json:"device_coap_port,omitempty"`

	// **参数说明**：设备接入COAPS协议端口, 默认值：5684
	DeviceCoapsPort *int32 `json:"device_coaps_port,omitempty"`

	// **参数说明**：设备接入MQTT协议端口, 默认值：1883
	DeviceMqttPort *int32 `json:"device_mqtt_port,omitempty"`

	// **参数说明**：设备接入MQTTS协议端口, 默认值：8883
	DeviceMqttsPort *int32 `json:"device_mqtts_port,omitempty"`

	// **参数说明**：设备接入HTTPS协议端口, 默认值：443
	DeviceHttpsPort *int32 `json:"device_https_port,omitempty"`
}

func (o Port) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Port struct{}"
	}

	return strings.Join([]string{"Port", string(data)}, " ")
}
