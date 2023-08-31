package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EquipmentLanItem struct {

	// 接口名字
	InterfaceName string `json:"interface_name"`

	// 接口类型
	InterfaceType EquipmentLanItemInterfaceType `json:"interface_type"`

	// VlanID
	VlanId *int32 `json:"vlan_id,omitempty"`

	// IPv4地址
	IpAddress string `json:"ip_address"`

	// DHCP开关
	Dhcp bool `json:"dhcp"`

	// DHCP地址池起始IP地址
	StartIpAddress *string `json:"start_ip_address,omitempty"`

	// DHCP地址池结束IP地址
	EndIpAddress *string `json:"end_ip_address,omitempty"`

	// 地址租期(分钟)
	LeaseTime *int32 `json:"lease_time,omitempty"`

	// 发布到企业连接网络
	PostToCloud bool `json:"post_to_cloud"`
}

func (o EquipmentLanItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentLanItem struct{}"
	}

	return strings.Join([]string{"EquipmentLanItem", string(data)}, " ")
}

type EquipmentLanItemInterfaceType struct {
	value string
}

type EquipmentLanItemInterfaceTypeEnum struct {
	L3_MAIN_INTERFACE EquipmentLanItemInterfaceType
	L3_SUB_INTERFACE  EquipmentLanItemInterfaceType
	VLAN_INTERFACE    EquipmentLanItemInterfaceType
}

func GetEquipmentLanItemInterfaceTypeEnum() EquipmentLanItemInterfaceTypeEnum {
	return EquipmentLanItemInterfaceTypeEnum{
		L3_MAIN_INTERFACE: EquipmentLanItemInterfaceType{
			value: "l3 main interface",
		},
		L3_SUB_INTERFACE: EquipmentLanItemInterfaceType{
			value: "l3 sub interface",
		},
		VLAN_INTERFACE: EquipmentLanItemInterfaceType{
			value: "vlan interface",
		},
	}
}

func (c EquipmentLanItemInterfaceType) Value() string {
	return c.value
}

func (c EquipmentLanItemInterfaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentLanItemInterfaceType) UnmarshalJSON(b []byte) error {
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
