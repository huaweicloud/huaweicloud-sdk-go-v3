package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type DevicesInGroup struct {
	// 设备ID

	DeviceId *int32 `json:"device_id,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 产品ID

	ProductId *int32 `json:"product_id,omitempty"`
	// 产品名称

	ProductName *string `json:"product_name,omitempty"`
	// 设备状态 0-启用 1-禁用

	Status *DevicesInGroupStatus `json:"status,omitempty"`
	// 是否在线 0-未连接 1-在线 2-离线

	OnlineStatus *DevicesInGroupOnlineStatus `json:"online_status,omitempty"`
}

func (o DevicesInGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicesInGroup struct{}"
	}

	return strings.Join([]string{"DevicesInGroup", string(data)}, " ")
}

type DevicesInGroupStatus struct {
	value int32
}

type DevicesInGroupStatusEnum struct {
	E_0 DevicesInGroupStatus
	E_1 DevicesInGroupStatus
}

func GetDevicesInGroupStatusEnum() DevicesInGroupStatusEnum {
	return DevicesInGroupStatusEnum{
		E_0: DevicesInGroupStatus{
			value: 0,
		}, E_1: DevicesInGroupStatus{
			value: 1,
		},
	}
}

func (c DevicesInGroupStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DevicesInGroupStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type DevicesInGroupOnlineStatus struct {
	value int32
}

type DevicesInGroupOnlineStatusEnum struct {
	E_0 DevicesInGroupOnlineStatus
	E_1 DevicesInGroupOnlineStatus
	E_2 DevicesInGroupOnlineStatus
}

func GetDevicesInGroupOnlineStatusEnum() DevicesInGroupOnlineStatusEnum {
	return DevicesInGroupOnlineStatusEnum{
		E_0: DevicesInGroupOnlineStatus{
			value: 0,
		}, E_1: DevicesInGroupOnlineStatus{
			value: 1,
		}, E_2: DevicesInGroupOnlineStatus{
			value: 2,
		},
	}
}

func (c DevicesInGroupOnlineStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DevicesInGroupOnlineStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
