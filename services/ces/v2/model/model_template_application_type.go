package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateApplicationType 模板应用类型。ALL_DIMENSION：所有维度, ONE_DIMENSION：同一维度。
type TemplateApplicationType struct {
	value string
}

type TemplateApplicationTypeEnum struct {
	ALL_DIMENSION TemplateApplicationType
	ONE_DIMENSION TemplateApplicationType
}

func GetTemplateApplicationTypeEnum() TemplateApplicationTypeEnum {
	return TemplateApplicationTypeEnum{
		ALL_DIMENSION: TemplateApplicationType{
			value: "ALL_DIMENSION",
		},
		ONE_DIMENSION: TemplateApplicationType{
			value: "ONE_DIMENSION",
		},
	}
}

func (c TemplateApplicationType) Value() string {
	return c.value
}

func (c TemplateApplicationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateApplicationType) UnmarshalJSON(b []byte) error {
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
