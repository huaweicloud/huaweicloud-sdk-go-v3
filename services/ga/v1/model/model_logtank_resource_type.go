package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LogtankResourceType 云日志的资源类型。 取值范围： LISTENER：监听器
type LogtankResourceType struct {
	value string
}

type LogtankResourceTypeEnum struct {
	LISTENER LogtankResourceType
}

func GetLogtankResourceTypeEnum() LogtankResourceTypeEnum {
	return LogtankResourceTypeEnum{
		LISTENER: LogtankResourceType{
			value: "LISTENER",
		},
	}
}

func (c LogtankResourceType) Value() string {
	return c.value
}

func (c LogtankResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LogtankResourceType) UnmarshalJSON(b []byte) error {
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
