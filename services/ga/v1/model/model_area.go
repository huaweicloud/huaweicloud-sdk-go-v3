package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Area 地区，取值： - OUTOFCM：中国大陆以外 - CM：中国大陆
type Area struct {
	value string
}

type AreaEnum struct {
	OUTOFCM Area
	CM      Area
}

func GetAreaEnum() AreaEnum {
	return AreaEnum{
		OUTOFCM: Area{
			value: "OUTOFCM",
		},
		CM: Area{
			value: "CM",
		},
	}
}

func (c Area) Value() string {
	return c.value
}

func (c Area) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Area) UnmarshalJSON(b []byte) error {
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
