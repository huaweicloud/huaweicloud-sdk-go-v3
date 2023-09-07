package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EnvironmentKindObj API类型，固定值“Environment”，该值不可修改。
type EnvironmentKindObj struct {
	value string
}

type EnvironmentKindObjEnum struct {
	ENVIRONMENT EnvironmentKindObj
}

func GetEnvironmentKindObjEnum() EnvironmentKindObjEnum {
	return EnvironmentKindObjEnum{
		ENVIRONMENT: EnvironmentKindObj{
			value: "Environment",
		},
	}
}

func (c EnvironmentKindObj) Value() string {
	return c.value
}

func (c EnvironmentKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentKindObj) UnmarshalJSON(b []byte) error {
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
