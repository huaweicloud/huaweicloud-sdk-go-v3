package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExternalAccessType 类型。
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

func (c ExternalAccessType) Value() string {
	return c.value
}

func (c ExternalAccessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExternalAccessType) UnmarshalJSON(b []byte) error {
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
