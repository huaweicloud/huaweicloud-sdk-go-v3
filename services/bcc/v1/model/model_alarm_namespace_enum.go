package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmNamespaceEnum 告警命名空间，取值范围：SYS.CBR,SYS.RDS,SYS.GaussDB
type AlarmNamespaceEnum struct {
	value string
}

type AlarmNamespaceEnumEnum struct {
	SYS_CBR      AlarmNamespaceEnum
	SYS_RDS      AlarmNamespaceEnum
	SYS_GAUSS_DB AlarmNamespaceEnum
}

func GetAlarmNamespaceEnumEnum() AlarmNamespaceEnumEnum {
	return AlarmNamespaceEnumEnum{
		SYS_CBR: AlarmNamespaceEnum{
			value: "SYS.CBR",
		},
		SYS_RDS: AlarmNamespaceEnum{
			value: "SYS.RDS",
		},
		SYS_GAUSS_DB: AlarmNamespaceEnum{
			value: "SYS.GaussDB",
		},
	}
}

func (c AlarmNamespaceEnum) Value() string {
	return c.value
}

func (c AlarmNamespaceEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmNamespaceEnum) UnmarshalJSON(b []byte) error {
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
