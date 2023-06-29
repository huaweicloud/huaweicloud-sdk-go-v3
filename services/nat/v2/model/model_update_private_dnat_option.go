package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePrivateDnatOption 更新DNAT规则的请求体。
type UpdatePrivateDnatOption struct {

	// DNAT规则的描述。
	Description *string `json:"description,omitempty"`

	// 中转IP的ID。
	TransitIpId *string `json:"transit_ip_id,omitempty"`

	// 网络接口ID，支持计算、ELB、VIP等实例的网络接口。
	NetworkInterfaceId *string `json:"network_interface_id,omitempty"`

	// 后端实例的私网IP地址。
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// 协议类型。 目前支持TCP/tcp、UDP/udp、ANY/any。 对应协议号6、17、0。
	Protocol *UpdatePrivateDnatOptionProtocol `json:"protocol,omitempty"`

	// 后端实例的端口号。
	InternalServicePort *string `json:"internal_service_port,omitempty"`

	// 中转IP的端口号。
	TransitServicePort *string `json:"transit_service_port,omitempty"`
}

func (o UpdatePrivateDnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateDnatOption struct{}"
	}

	return strings.Join([]string{"UpdatePrivateDnatOption", string(data)}, " ")
}

type UpdatePrivateDnatOptionProtocol struct {
	value string
}

type UpdatePrivateDnatOptionProtocolEnum struct {
	TCP UpdatePrivateDnatOptionProtocol
	UDP UpdatePrivateDnatOptionProtocol
	ANY UpdatePrivateDnatOptionProtocol
}

func GetUpdatePrivateDnatOptionProtocolEnum() UpdatePrivateDnatOptionProtocolEnum {
	return UpdatePrivateDnatOptionProtocolEnum{
		TCP: UpdatePrivateDnatOptionProtocol{
			value: "tcp",
		},
		UDP: UpdatePrivateDnatOptionProtocol{
			value: "udp",
		},
		ANY: UpdatePrivateDnatOptionProtocol{
			value: "any",
		},
	}
}

func (c UpdatePrivateDnatOptionProtocol) Value() string {
	return c.value
}

func (c UpdatePrivateDnatOptionProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePrivateDnatOptionProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
