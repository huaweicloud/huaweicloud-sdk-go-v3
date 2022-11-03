package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 地区，取值： - OUTOFCM：中国大陆以外 - CM：中国大陆
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
