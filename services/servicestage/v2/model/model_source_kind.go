package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceKind 来源类型。支持源码code和artifact软件包。
type SourceKind struct {
	value string
}

type SourceKindEnum struct {
	CODE     SourceKind
	ARTIFACT SourceKind
}

func GetSourceKindEnum() SourceKindEnum {
	return SourceKindEnum{
		CODE: SourceKind{
			value: "code",
		},
		ARTIFACT: SourceKind{
			value: "artifact",
		},
	}
}

func (c SourceKind) Value() string {
	return c.value
}

func (c SourceKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceKind) UnmarshalJSON(b []byte) error {
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
