package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceRestartRequsetBody struct {
	// 在线调试时必填。

	Restart *InstanceRestartRequsetBodyRestart `json:"restart,omitempty"`
}

func (o InstanceRestartRequsetBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InstanceRestartRequsetBody struct{}"
	}

	return strings.Join([]string{"InstanceRestartRequsetBody", string(data)}, " ")
}

type InstanceRestartRequsetBodyRestart struct {
	value string
}

type InstanceRestartRequsetBodyRestartEnum struct {
	TRUE InstanceRestartRequsetBodyRestart
}

func GetInstanceRestartRequsetBodyRestartEnum() InstanceRestartRequsetBodyRestartEnum {
	return InstanceRestartRequsetBodyRestartEnum{
		TRUE: InstanceRestartRequsetBodyRestart{
			value: "true",
		},
	}
}

func (c InstanceRestartRequsetBodyRestart) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *InstanceRestartRequsetBodyRestart) UnmarshalJSON(b []byte) error {
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
