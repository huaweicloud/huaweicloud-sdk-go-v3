package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LtsLogType LTS日志类型，不能为空，具有以下日志类型。 - audit_log
type LtsLogType struct {
	value string
}

type LtsLogTypeEnum struct {
	AUDIT_LOG LtsLogType
}

func GetLtsLogTypeEnum() LtsLogTypeEnum {
	return LtsLogTypeEnum{
		AUDIT_LOG: LtsLogType{
			value: "audit_log",
		},
	}
}

func (c LtsLogType) Value() string {
	return c.value
}

func (c LtsLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LtsLogType) UnmarshalJSON(b []byte) error {
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
