package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 包周期付款类型 - FULL: 全预付 - HALF：半预付
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
