package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NonRequiredGcbChargeMode 全域互联带宽计费类型。
type NonRequiredGcbChargeMode struct {

	// 功能说明：描述计费类型，描述可选计费类型。默认开放按带宽计费，传统95计费租户白名单控制。 取值范围：     bwd: 按带宽计费     95: 按传统型95计费     95avr: 按传统型日95计费
	ChargeMode *NonRequiredGcbChargeModeChargeMode `json:"charge_mode,omitempty"`
}

func (o NonRequiredGcbChargeMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonRequiredGcbChargeMode struct{}"
	}

	return strings.Join([]string{"NonRequiredGcbChargeMode", string(data)}, " ")
}

type NonRequiredGcbChargeModeChargeMode struct {
	value string
}

type NonRequiredGcbChargeModeChargeModeEnum struct {
	BWD     NonRequiredGcbChargeModeChargeMode
	E_95    NonRequiredGcbChargeModeChargeMode
	E_95AVR NonRequiredGcbChargeModeChargeMode
}

func GetNonRequiredGcbChargeModeChargeModeEnum() NonRequiredGcbChargeModeChargeModeEnum {
	return NonRequiredGcbChargeModeChargeModeEnum{
		BWD: NonRequiredGcbChargeModeChargeMode{
			value: "bwd",
		},
		E_95: NonRequiredGcbChargeModeChargeMode{
			value: "95",
		},
		E_95AVR: NonRequiredGcbChargeModeChargeMode{
			value: "95avr",
		},
	}
}

func (c NonRequiredGcbChargeModeChargeMode) Value() string {
	return c.value
}

func (c NonRequiredGcbChargeModeChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NonRequiredGcbChargeModeChargeMode) UnmarshalJSON(b []byte) error {
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
