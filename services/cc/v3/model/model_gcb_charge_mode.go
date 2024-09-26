package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GcbChargeMode 全域互联带宽计费类型。
type GcbChargeMode struct {

	// 功能说明：描述计费类型，描述可选计费类型。默认开放按带宽计费，传统95计费租户白名单控制。 取值范围：     bwd: 按带宽计费     95: 按传统型95计费     95avr: 按传统型日95计费
	ChargeMode GcbChargeModeChargeMode `json:"charge_mode"`
}

func (o GcbChargeMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbChargeMode struct{}"
	}

	return strings.Join([]string{"GcbChargeMode", string(data)}, " ")
}

type GcbChargeModeChargeMode struct {
	value string
}

type GcbChargeModeChargeModeEnum struct {
	BWD     GcbChargeModeChargeMode
	E_95    GcbChargeModeChargeMode
	E_95AVR GcbChargeModeChargeMode
}

func GetGcbChargeModeChargeModeEnum() GcbChargeModeChargeModeEnum {
	return GcbChargeModeChargeModeEnum{
		BWD: GcbChargeModeChargeMode{
			value: "bwd",
		},
		E_95: GcbChargeModeChargeMode{
			value: "95",
		},
		E_95AVR: GcbChargeModeChargeMode{
			value: "95avr",
		},
	}
}

func (c GcbChargeModeChargeMode) Value() string {
	return c.value
}

func (c GcbChargeModeChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GcbChargeModeChargeMode) UnmarshalJSON(b []byte) error {
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
