package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChargeMode **参数解释**：计费模式。 **取值范围**：   - postPaid：按需计费。   - prePaid：包周期计费。
type ChargeMode struct {
	value string
}

type ChargeModeEnum struct {
	PRE_PAID  ChargeMode
	POST_PAID ChargeMode
}

func GetChargeModeEnum() ChargeModeEnum {
	return ChargeModeEnum{
		PRE_PAID: ChargeMode{
			value: "prePaid",
		},
		POST_PAID: ChargeMode{
			value: "postPaid",
		},
	}
}

func (c ChargeMode) Value() string {
	return c.value
}

func (c ChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChargeMode) UnmarshalJSON(b []byte) error {
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
