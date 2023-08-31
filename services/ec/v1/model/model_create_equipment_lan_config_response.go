package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateEquipmentLanConfigResponse Response Object
type CreateEquipmentLanConfigResponse struct {

	// 接口名字
	InterfaceName *string `json:"interface_name,omitempty"`

	// 接口类型
	InterfaceType *CreateEquipmentLanConfigResponseInterfaceType `json:"interface_type,omitempty"`

	// VlanID
	VlanId *int32 `json:"vlan_id,omitempty"`

	// IPv4地址
	IpAddress *string `json:"ip_address,omitempty"`

	// DHCP开关
	Dhcp *bool `json:"dhcp,omitempty"`

	// DHCP地址池起始IP地址
	StartIpAddress *string `json:"start_ip_address,omitempty"`

	// DHCP地址池结束IP地址
	EndIpAddress *string `json:"end_ip_address,omitempty"`

	// 地址租期(分钟)
	LeaseTime *int32 `json:"lease_time,omitempty"`

	// 发布到企业连接网络
	PostToCloud    *bool `json:"post_to_cloud,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateEquipmentLanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEquipmentLanConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateEquipmentLanConfigResponse", string(data)}, " ")
}

type CreateEquipmentLanConfigResponseInterfaceType struct {
	value string
}

type CreateEquipmentLanConfigResponseInterfaceTypeEnum struct {
	L3_MAIN_INTERFACE CreateEquipmentLanConfigResponseInterfaceType
	L3_SUB_INTERFACE  CreateEquipmentLanConfigResponseInterfaceType
	VLAN_INTERFACE    CreateEquipmentLanConfigResponseInterfaceType
}

func GetCreateEquipmentLanConfigResponseInterfaceTypeEnum() CreateEquipmentLanConfigResponseInterfaceTypeEnum {
	return CreateEquipmentLanConfigResponseInterfaceTypeEnum{
		L3_MAIN_INTERFACE: CreateEquipmentLanConfigResponseInterfaceType{
			value: "l3 main interface",
		},
		L3_SUB_INTERFACE: CreateEquipmentLanConfigResponseInterfaceType{
			value: "l3 sub interface",
		},
		VLAN_INTERFACE: CreateEquipmentLanConfigResponseInterfaceType{
			value: "vlan interface",
		},
	}
}

func (c CreateEquipmentLanConfigResponseInterfaceType) Value() string {
	return c.value
}

func (c CreateEquipmentLanConfigResponseInterfaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEquipmentLanConfigResponseInterfaceType) UnmarshalJSON(b []byte) error {
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
