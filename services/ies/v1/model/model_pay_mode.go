package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PayMode 包周期付款类型 - FULL: 全预付 - HALF：半预付
type PayMode struct {
	value string
}

type PayModeEnum struct {
	FULL PayMode
	HALF PayMode
}

func GetPayModeEnum() PayModeEnum {
	return PayModeEnum{
		FULL: PayMode{
			value: "FULL",
		},
		HALF: PayMode{
			value: "HALF",
		},
	}
}

func (c PayMode) Value() string {
	return c.value
}

func (c PayMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PayMode) UnmarshalJSON(b []byte) error {
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
