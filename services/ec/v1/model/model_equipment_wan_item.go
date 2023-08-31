package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EquipmentWanItem struct {

	// 接口名字
	InterfaceName string `json:"interface_name"`

	// IP类型
	IpType *EquipmentWanItemIpType `json:"ip_type,omitempty"`

	// IPv4地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 网关IP
	GatewayIpAddress *string `json:"gateway_ip_address,omitempty"`

	// 最大上行带宽
	UplinkBandwidthSize *int32 `json:"uplink_bandwidth_size,omitempty"`

	// 优先级
	Priority EquipmentWanItemPriority `json:"priority"`

	// 是否使能本地上网
	NatOutbound *bool `json:"nat_outbound,omitempty"`

	// 最大上云带宽
	TunnelBandwidthSize *int32 `json:"tunnel_bandwidth_size,omitempty"`
}

func (o EquipmentWanItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentWanItem struct{}"
	}

	return strings.Join([]string{"EquipmentWanItem", string(data)}, " ")
}

type EquipmentWanItemIpType struct {
	value string
}

type EquipmentWanItemIpTypeEnum struct {
	STATIC EquipmentWanItemIpType
	DHCP   EquipmentWanItemIpType
}

func GetEquipmentWanItemIpTypeEnum() EquipmentWanItemIpTypeEnum {
	return EquipmentWanItemIpTypeEnum{
		STATIC: EquipmentWanItemIpType{
			value: "static",
		},
		DHCP: EquipmentWanItemIpType{
			value: "DHCP",
		},
	}
}

func (c EquipmentWanItemIpType) Value() string {
	return c.value
}

func (c EquipmentWanItemIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentWanItemIpType) UnmarshalJSON(b []byte) error {
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

type EquipmentWanItemPriority struct {
	value string
}

type EquipmentWanItemPriorityEnum struct {
	ACTIVE  EquipmentWanItemPriority
	STANDBY EquipmentWanItemPriority
}

func GetEquipmentWanItemPriorityEnum() EquipmentWanItemPriorityEnum {
	return EquipmentWanItemPriorityEnum{
		ACTIVE: EquipmentWanItemPriority{
			value: "Active",
		},
		STANDBY: EquipmentWanItemPriority{
			value: "Standby",
		},
	}
}

func (c EquipmentWanItemPriority) Value() string {
	return c.value
}

func (c EquipmentWanItemPriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentWanItemPriority) UnmarshalJSON(b []byte) error {
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
