package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentCategory 应用组件类型包括：Webapp、MicroService、Common。
type ComponentCategory struct {
	value string
}

type ComponentCategoryEnum struct {
	WEBAPP        ComponentCategory
	MICRO_SERVICE ComponentCategory
	COMMON        ComponentCategory
}

func GetComponentCategoryEnum() ComponentCategoryEnum {
	return ComponentCategoryEnum{
		WEBAPP: ComponentCategory{
			value: "Webapp",
		},
		MICRO_SERVICE: ComponentCategory{
			value: "MicroService",
		},
		COMMON: ComponentCategory{
			value: "Common",
		},
	}
}

func (c ComponentCategory) Value() string {
	return c.value
}

func (c ComponentCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentCategory) UnmarshalJSON(b []byte) error {
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
