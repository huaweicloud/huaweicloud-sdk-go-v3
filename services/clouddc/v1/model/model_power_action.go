package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PowerAction 电源状态： - power on：正常开启电源 - power off：正常关闭电源 - rebooting：重启
type PowerAction struct {

	// 电源操作
	Action PowerActionAction `json:"action"`

	// server id set
	ServerIdSet *[]string `json:"server_id_set,omitempty"`
}

func (o PowerAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PowerAction struct{}"
	}

	return strings.Join([]string{"PowerAction", string(data)}, " ")
}

type PowerActionAction struct {
	value string
}

type PowerActionActionEnum struct {
	POWER_ON  PowerActionAction
	POWER_OFF PowerActionAction
	REBOOTING PowerActionAction
}

func GetPowerActionActionEnum() PowerActionActionEnum {
	return PowerActionActionEnum{
		POWER_ON: PowerActionAction{
			value: "power on",
		},
		POWER_OFF: PowerActionAction{
			value: "power off",
		},
		REBOOTING: PowerActionAction{
			value: "rebooting",
		},
	}
}

func (c PowerActionAction) Value() string {
	return c.value
}

func (c PowerActionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PowerActionAction) UnmarshalJSON(b []byte) error {
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
