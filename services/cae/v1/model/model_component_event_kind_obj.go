package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentEventKindObj API类型，固定值“ComponentEvent”，该值不可修改。
type ComponentEventKindObj struct {
	value string
}

type ComponentEventKindObjEnum struct {
	COMPONENT_EVENT ComponentEventKindObj
}

func GetComponentEventKindObjEnum() ComponentEventKindObjEnum {
	return ComponentEventKindObjEnum{
		COMPONENT_EVENT: ComponentEventKindObj{
			value: "ComponentEvent",
		},
	}
}

func (c ComponentEventKindObj) Value() string {
	return c.value
}

func (c ComponentEventKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentEventKindObj) UnmarshalJSON(b []byte) error {
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
