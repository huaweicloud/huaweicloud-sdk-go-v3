package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessInfo 设备接入实例的接入信息。用户可以使用该结构体中的信息将应用服务器和设备接入到物联网平台。
type AccessInfo struct {

	// **参数说明**：接入地址的类型，如应用接入的HTTPS协议的取值为：APP_HTTPS，设备接入的MQTT协议的取值为：DEVICE_MQTT **取值范围**： - APP_HTTPS：应用接入HTTPS协议 - APP_AMQP：应用接入AMQP协议 - APP_MQTT：应用接入MQTT协议 - DEVICE_COAP：设备接入COAP协议 - DEVICE_MQTT：设备接入MQTT协议 - DEVICE_HTTPS：设备接入HTTPS协议
	Type *string `json:"type,omitempty"`

	// **参数说明**：实例的应用/设备的安全接入端口
	Port *int32 `json:"port,omitempty"`

	// **参数说明**：实例的应用/设备的非安全接入端口。返回null时表示该类型的接入地址不支持非安全端口接入。
	NonTlsPort *int32 `json:"non_tls_port,omitempty"`

	// **参数说明**：基于WebSocket的MQTT接入端口。返回null时表示该类型的接入地址不支持WebSocket端口接入。
	WebsocketPort *int32 `json:"websocket_port,omitempty"`

	// **参数说明**：实例的接入域名
	DomainName *string `json:"domain_name,omitempty"`

	// **参数说明**：实例的私网接入地址列表
	PrivateAddresses *[]string `json:"private_addresses,omitempty"`

	// **参数说明**：实例的公网接入地址
	PublicAddress *[]string `json:"public_address,omitempty"`

	// **参数说明**：实例的ipv6接入地址列表
	Ipv6Address *[]string `json:"ipv6_address,omitempty"`

	IpWhitelist *IpWhiteList `json:"ip_whitelist,omitempty"`
}

func (o AccessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessInfo struct{}"
	}

	return strings.Join([]string{"AccessInfo", string(data)}, " ")
}
