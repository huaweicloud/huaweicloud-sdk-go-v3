package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建DNAT规则的请求体。
type CreatePrivateDnatOption struct {

	// DNAT规则的描述。
	Description *string `json:"description,omitempty"`

	// 中转IP的ID。
	TransitIpId string `json:"transit_ip_id"`

	// 网络接口ID，支持计算、ELB、VIP等实例的网络接口。
	NetworkInterfaceId *string `json:"network_interface_id,omitempty"`

	// 私网NAT网关实例的ID。
	GatewayId string `json:"gateway_id"`

	// 协议类型。 目前支持TCP/tcp、UDP/udp、ANY/any。 对应协议号6、17、0。
	Protocol *CreatePrivateDnatOptionProtocol `json:"protocol,omitempty"`

	// 后端实例的私网IP地址。
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// 后端实例的端口号。
	InternalServicePort *string `json:"internal_service_port,omitempty"`

	// 中转IP的端口号。
	TransitServicePort *string `json:"transit_service_port,omitempty"`
}

func (o CreatePrivateDnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateDnatOption struct{}"
	}

	return strings.Join([]string{"CreatePrivateDnatOption", string(data)}, " ")
}

type CreatePrivateDnatOptionProtocol struct {
	value string
}

type CreatePrivateDnatOptionProtocolEnum struct {
	TCP CreatePrivateDnatOptionProtocol
	UDP CreatePrivateDnatOptionProtocol
	ANY CreatePrivateDnatOptionProtocol
}

func GetCreatePrivateDnatOptionProtocolEnum() CreatePrivateDnatOptionProtocolEnum {
	return CreatePrivateDnatOptionProtocolEnum{
		TCP: CreatePrivateDnatOptionProtocol{
			value: "tcp",
		},
		UDP: CreatePrivateDnatOptionProtocol{
			value: "udp",
		},
		ANY: CreatePrivateDnatOptionProtocol{
			value: "any",
		},
	}
}

func (c CreatePrivateDnatOptionProtocol) Value() string {
	return c.value
}

func (c CreatePrivateDnatOptionProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivateDnatOptionProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
