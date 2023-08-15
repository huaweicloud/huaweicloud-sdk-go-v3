package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ActionType struct {
	value string
}

type ActionTypeEnum struct {
	START ActionType
	STOP  ActionType
}

func GetActionTypeEnum() ActionTypeEnum {
	return ActionTypeEnum{
		START: ActionType{
			value: "start",
		},
		STOP: ActionType{
			value: "stop",
		},
	}
}

func (c ActionType) Value() string {
	return c.value
}

func (c ActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionType) UnmarshalJSON(b []byte) error {
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
