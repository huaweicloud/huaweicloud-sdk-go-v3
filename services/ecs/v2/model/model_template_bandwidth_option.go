package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TemplateBandwidthOption struct {

	// 带宽类型
	ShareType *TemplateBandwidthOptionShareType `json:"share_type,omitempty"`

	// 带宽大小
	Size *int32 `json:"size,omitempty"`

	// 计费类型
	ChargeMode *TemplateBandwidthOptionChargeMode `json:"charge_mode,omitempty"`

	// 带宽ID，创建WHOLE类型带宽的弹性IP时可以指定之前的共享带宽创建
	Id *string `json:"id,omitempty"`
}

func (o TemplateBandwidthOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateBandwidthOption struct{}"
	}

	return strings.Join([]string{"TemplateBandwidthOption", string(data)}, " ")
}

type TemplateBandwidthOptionShareType struct {
	value string
}

type TemplateBandwidthOptionShareTypeEnum struct {
	PER   TemplateBandwidthOptionShareType
	WHOLE TemplateBandwidthOptionShareType
}

func GetTemplateBandwidthOptionShareTypeEnum() TemplateBandwidthOptionShareTypeEnum {
	return TemplateBandwidthOptionShareTypeEnum{
		PER: TemplateBandwidthOptionShareType{
			value: "PER",
		},
		WHOLE: TemplateBandwidthOptionShareType{
			value: "WHOLE",
		},
	}
}

func (c TemplateBandwidthOptionShareType) Value() string {
	return c.value
}

func (c TemplateBandwidthOptionShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateBandwidthOptionShareType) UnmarshalJSON(b []byte) error {
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

type TemplateBandwidthOptionChargeMode struct {
	value string
}

type TemplateBandwidthOptionChargeModeEnum struct {
	BANDWIDTH TemplateBandwidthOptionChargeMode
	TRAFFIC   TemplateBandwidthOptionChargeMode
}

func GetTemplateBandwidthOptionChargeModeEnum() TemplateBandwidthOptionChargeModeEnum {
	return TemplateBandwidthOptionChargeModeEnum{
		BANDWIDTH: TemplateBandwidthOptionChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: TemplateBandwidthOptionChargeMode{
			value: "traffic",
		},
	}
}

func (c TemplateBandwidthOptionChargeMode) Value() string {
	return c.value
}

func (c TemplateBandwidthOptionChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateBandwidthOptionChargeMode) UnmarshalJSON(b []byte) error {
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
