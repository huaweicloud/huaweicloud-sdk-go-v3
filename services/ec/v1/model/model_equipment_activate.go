package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EquipmentActivate struct {

	// 设备名称
	Name string `json:"name"`

	// 设备类型
	Type EquipmentActivateType `json:"type"`

	// esn
	Esn string `json:"esn"`

	// 高可用类型
	HaType EquipmentActivateHaType `json:"ha_type"`
}

func (o EquipmentActivate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentActivate struct{}"
	}

	return strings.Join([]string{"EquipmentActivate", string(data)}, " ")
}

type EquipmentActivateType struct {
	value string
}

type EquipmentActivateTypeEnum struct {
	STANDARD EquipmentActivateType
	SOHO     EquipmentActivateType
}

func GetEquipmentActivateTypeEnum() EquipmentActivateTypeEnum {
	return EquipmentActivateTypeEnum{
		STANDARD: EquipmentActivateType{
			value: "standard",
		},
		SOHO: EquipmentActivateType{
			value: "soho",
		},
	}
}

func (c EquipmentActivateType) Value() string {
	return c.value
}

func (c EquipmentActivateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentActivateType) UnmarshalJSON(b []byte) error {
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

type EquipmentActivateHaType struct {
	value string
}

type EquipmentActivateHaTypeEnum struct {
	ACTIVE  EquipmentActivateHaType
	STANDBY EquipmentActivateHaType
}

func GetEquipmentActivateHaTypeEnum() EquipmentActivateHaTypeEnum {
	return EquipmentActivateHaTypeEnum{
		ACTIVE: EquipmentActivateHaType{
			value: "Active",
		},
		STANDBY: EquipmentActivateHaType{
			value: "Standby",
		},
	}
}

func (c EquipmentActivateHaType) Value() string {
	return c.value
}

func (c EquipmentActivateHaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentActivateHaType) UnmarshalJSON(b []byte) error {
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
