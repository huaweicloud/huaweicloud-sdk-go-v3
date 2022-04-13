package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 类型。
type ExternalAccessType struct {
	value string
}

type ExternalAccessTypeEnum struct {
	AUTO_GENERATED ExternalAccessType
	SPECIFIED      ExternalAccessType
	IP_ADDR        ExternalAccessType
}

func GetExternalAccessTypeEnum() ExternalAccessTypeEnum {
	return ExternalAccessTypeEnum{
		AUTO_GENERATED: ExternalAccessType{
			value: "AUTO_GENERATED",
		},
		SPECIFIED: ExternalAccessType{
			value: "SPECIFIED",
		},
		IP_ADDR: ExternalAccessType{
			value: "IP_ADDR",
		},
	}
}

func (c ExternalAccessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExternalAccessType) UnmarshalJSON(b []byte) error {
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
