package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// PrivateDnat DNAT规则的响应体。
type PrivateDnat struct {

	// DNAT规则的ID。
	Id *string `json:"id,omitempty"`

	// 项目的ID。
	ProjectId *string `json:"project_id,omitempty"`

	// DNAT规则的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 中转IP的ID。
	TransitIpId *string `json:"transit_ip_id,omitempty"`

	// 私网NAT网关实例的ID。
	GatewayId *string `json:"gateway_id,omitempty"`

	// 网络接口ID，支持计算、ELBV2、ELBV3、VIP等实例的端口。
	NetworkInterfaceId *string `json:"network_interface_id,omitempty"`

	// DNAT规则后端的类型。 取值： - COMPUTE：后端为计算实例。 - VIP：后端为VIP的实例。 - ELB：后端为ELBv2的实例。 - ELBv3：后端为ELBv3的实例。 - CUSTOMIZE：后端为自定义IP。
	Type *string `json:"type,omitempty"`

	// 协议类型。 目前支持TCP/tcp、UDP/udp、ANY/any。 对应协议号6、17、0。
	Protocol *PrivateDnatProtocol `json:"protocol,omitempty"`

	// 后端资源（计算实例、ELBV2、ELBV3、VIP等）的私网IP地址。
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// 后端实例的端口号（计算实例、ELBV2、ELBV3、VIP等)。
	InternalServicePort *string `json:"internal_service_port,omitempty"`

	// 中转IP的端口号。
	TransitServicePort *string `json:"transit_service_port,omitempty"`

	// 企业项目ID。创建DNAT规则时，关联的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// DNAT规则的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// DNAT规则的更新时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 私网NAT的DNAT规则状态。 取值为： - ACTIVE：正常运行 - FROZEN：冻结 - INACTIVE：不可用
	Status *PrivateDnatStatus `json:"status,omitempty"`
}

func (o PrivateDnat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateDnat struct{}"
	}

	return strings.Join([]string{"PrivateDnat", string(data)}, " ")
}

type PrivateDnatProtocol struct {
	value string
}

type PrivateDnatProtocolEnum struct {
	TCP PrivateDnatProtocol
	UDP PrivateDnatProtocol
	ANY PrivateDnatProtocol
}

func GetPrivateDnatProtocolEnum() PrivateDnatProtocolEnum {
	return PrivateDnatProtocolEnum{
		TCP: PrivateDnatProtocol{
			value: "tcp",
		},
		UDP: PrivateDnatProtocol{
			value: "udp",
		},
		ANY: PrivateDnatProtocol{
			value: "any",
		},
	}
}

func (c PrivateDnatProtocol) Value() string {
	return c.value
}

func (c PrivateDnatProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateDnatProtocol) UnmarshalJSON(b []byte) error {
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

type PrivateDnatStatus struct {
	value string
}

type PrivateDnatStatusEnum struct {
	ACTIVE   PrivateDnatStatus
	FROZEN   PrivateDnatStatus
	INACTIVE PrivateDnatStatus
}

func GetPrivateDnatStatusEnum() PrivateDnatStatusEnum {
	return PrivateDnatStatusEnum{
		ACTIVE: PrivateDnatStatus{
			value: "ACTIVE",
		},
		FROZEN: PrivateDnatStatus{
			value: "FROZEN",
		},
		INACTIVE: PrivateDnatStatus{
			value: "INACTIVE",
		},
	}
}

func (c PrivateDnatStatus) Value() string {
	return c.value
}

func (c PrivateDnatStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateDnatStatus) UnmarshalJSON(b []byte) error {
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
