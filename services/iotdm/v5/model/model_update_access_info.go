package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessInfo 设备接入实例的接入信息。用户可以使用该结构体中的信息将应用服务器和设备接入到物联网平台。 约束：只有企业版实例支持自定义接入信息。
type UpdateAccessInfo struct {

	// **参数说明**：接入地址的类型，如应用接入的HTTPS协议的取值为：APP_HTTPS，设备接入的MQTT协议的取值为：DEVICE_MQTT。 **取值范围**： - APP_HTTPS：应用接入HTTPS协议 - APP_AMQP：应用接入AMQP协议 - APP_MQTT：应用接入MQTT协议 - DEVICE_COAP：设备接入COAP协议 - DEVICE_MQTT：设备接入MQTT协议 - DEVICE_HTTPS：设备接入HTTPS协议
	AccessType string `json:"access_type"`

	// **参数说明**：接入域名，如果需要更新域名，则携带该字段。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数说明**：是否配置公网接入地址。约束：access_type为APP_HTTPS/APP_AMQP/APP_MQTT的公网地址会同时被解绑或绑定。access_type为DEVICE_MQTT/DEVICE_HTTPS的公网地址会同时被解绑或绑定。 **取值范围**： - true：配置公网接入地址，平台将自动分配公网接入地址。 - false: 解绑公网接入地址，解绑公网地址不会被删除，再次配置公网地址将绑定原来的公网地址。
	PublicAddressesEnable *bool `json:"public_addresses_enable,omitempty"`

	IpWhitelist *IpWhiteList `json:"ip_whitelist,omitempty"`
}

func (o UpdateAccessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessInfo struct{}"
	}

	return strings.Join([]string{"UpdateAccessInfo", string(data)}, " ")
}
