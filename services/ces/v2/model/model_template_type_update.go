package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateTypeUpdate **参数解释**： 自定义告警模板类型 **约束限制**： 不涉及。 **取值范围**： 0：指标；2： 事件。 **默认取值**： 0
type TemplateTypeUpdate struct {
	value int32
}

type TemplateTypeUpdateEnum struct {
	E_0 TemplateTypeUpdate
	E_2 TemplateTypeUpdate
}

func GetTemplateTypeUpdateEnum() TemplateTypeUpdateEnum {
	return TemplateTypeUpdateEnum{
		E_0: TemplateTypeUpdate{
			value: 0,
		}, E_2: TemplateTypeUpdate{
			value: 2,
		},
	}
}

func (c TemplateTypeUpdate) Value() int32 {
	return c.value
}

func (c TemplateTypeUpdate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateTypeUpdate) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
