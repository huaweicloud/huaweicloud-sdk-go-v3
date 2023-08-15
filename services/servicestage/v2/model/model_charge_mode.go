package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChargeMode 收费模式，支持provided、on_demanded、monthly三种模式。  默认provided，表示使用用户提供的已有资源，无需收费。
type ChargeMode struct {
	value string
}

type ChargeModeEnum struct {
	PROVIDED    ChargeMode
	ON_DEMANDED ChargeMode
	MONTHLY     ChargeMode
}

func GetChargeModeEnum() ChargeModeEnum {
	return ChargeModeEnum{
		PROVIDED: ChargeMode{
			value: "provided",
		},
		ON_DEMANDED: ChargeMode{
			value: "on_demanded",
		},
		MONTHLY: ChargeMode{
			value: "monthly",
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
