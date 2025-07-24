package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// State 硬件启用状态： Enabled：启用 Unenabled：未启用 Unknown：未知
type State struct {
	value string
}

type StateEnum struct {
	ENABLED   State
	UNENABLED State
	UNKNOWN   State
}

func GetStateEnum() StateEnum {
	return StateEnum{
		ENABLED: State{
			value: "Enabled",
		},
		UNENABLED: State{
			value: "Unenabled",
		},
		UNKNOWN: State{
			value: "Unknown",
		},
	}
}

func (c State) Value() string {
	return c.value
}

func (c State) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *State) UnmarshalJSON(b []byte) error {
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
