package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentConfigurationKindObj API类型，固定值“ComponentConfiguration”，该值不可修改。
type ComponentConfigurationKindObj struct {
	value string
}

type ComponentConfigurationKindObjEnum struct {
	COMPONENT_CONFIGURATION ComponentConfigurationKindObj
}

func GetComponentConfigurationKindObjEnum() ComponentConfigurationKindObjEnum {
	return ComponentConfigurationKindObjEnum{
		COMPONENT_CONFIGURATION: ComponentConfigurationKindObj{
			value: "ComponentConfiguration",
		},
	}
}

func (c ComponentConfigurationKindObj) Value() string {
	return c.value
}

func (c ComponentConfigurationKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentConfigurationKindObj) UnmarshalJSON(b []byte) error {
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
