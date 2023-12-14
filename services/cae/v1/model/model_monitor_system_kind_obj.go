package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MonitorSystemKindObj API类型，固定值“MonitorSystem”，该值不可修改。
type MonitorSystemKindObj struct {
	value string
}

type MonitorSystemKindObjEnum struct {
	MONITOR_SYSTEM MonitorSystemKindObj
}

func GetMonitorSystemKindObjEnum() MonitorSystemKindObjEnum {
	return MonitorSystemKindObjEnum{
		MONITOR_SYSTEM: MonitorSystemKindObj{
			value: "MonitorSystem",
		},
	}
}

func (c MonitorSystemKindObj) Value() string {
	return c.value
}

func (c MonitorSystemKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MonitorSystemKindObj) UnmarshalJSON(b []byte) error {
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
