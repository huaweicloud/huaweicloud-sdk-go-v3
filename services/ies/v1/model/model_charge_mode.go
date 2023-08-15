package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChargeMode 计费模式。 - prepaid：包年/包月
type ChargeMode struct {
	value string
}

type ChargeModeEnum struct {
	PREPAID ChargeMode
}

func GetChargeModeEnum() ChargeModeEnum {
	return ChargeModeEnum{
		PREPAID: ChargeMode{
			value: "prepaid",
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
