package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentKindObj API类型，固定值“Component”，该值不可修改。
type ComponentKindObj struct {
	value string
}

type ComponentKindObjEnum struct {
	COMPONENT ComponentKindObj
}

func GetComponentKindObjEnum() ComponentKindObjEnum {
	return ComponentKindObjEnum{
		COMPONENT: ComponentKindObj{
			value: "Component",
		},
	}
}

func (c ComponentKindObj) Value() string {
	return c.value
}

func (c ComponentKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentKindObj) UnmarshalJSON(b []byte) error {
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
