package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEquipmentLanConfigResponse Response Object
type UpdateEquipmentLanConfigResponse struct {

	// 接口名字
	InterfaceName *string `json:"interface_name,omitempty"`

	// 接口类型
	InterfaceType *UpdateEquipmentLanConfigResponseInterfaceType `json:"interface_type,omitempty"`

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

func (o UpdateEquipmentLanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentLanConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentLanConfigResponse", string(data)}, " ")
}

type UpdateEquipmentLanConfigResponseInterfaceType struct {
	value string
}

type UpdateEquipmentLanConfigResponseInterfaceTypeEnum struct {
	L3_MAIN_INTERFACE UpdateEquipmentLanConfigResponseInterfaceType
	L3_SUB_INTERFACE  UpdateEquipmentLanConfigResponseInterfaceType
	VLAN_INTERFACE    UpdateEquipmentLanConfigResponseInterfaceType
}

func GetUpdateEquipmentLanConfigResponseInterfaceTypeEnum() UpdateEquipmentLanConfigResponseInterfaceTypeEnum {
	return UpdateEquipmentLanConfigResponseInterfaceTypeEnum{
		L3_MAIN_INTERFACE: UpdateEquipmentLanConfigResponseInterfaceType{
			value: "l3 main interface",
		},
		L3_SUB_INTERFACE: UpdateEquipmentLanConfigResponseInterfaceType{
			value: "l3 sub interface",
		},
		VLAN_INTERFACE: UpdateEquipmentLanConfigResponseInterfaceType{
			value: "vlan interface",
		},
	}
}

func (c UpdateEquipmentLanConfigResponseInterfaceType) Value() string {
	return c.value
}

func (c UpdateEquipmentLanConfigResponseInterfaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEquipmentLanConfigResponseInterfaceType) UnmarshalJSON(b []byte) error {
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
