package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DatabaseStatusEnum struct {
	value string
}

type DatabaseStatusEnumEnum struct {
	DEPLOYING DatabaseStatusEnum
	NORMAL    DatabaseStatusEnum
	UPDATINTG DatabaseStatusEnum
	ABNORMAL  DatabaseStatusEnum
	FREEZE    DatabaseStatusEnum
	DELETING  DatabaseStatusEnum
}

func GetDatabaseStatusEnumEnum() DatabaseStatusEnumEnum {
	return DatabaseStatusEnumEnum{
		DEPLOYING: DatabaseStatusEnum{
			value: "DEPLOYING",
		},
		NORMAL: DatabaseStatusEnum{
			value: "NORMAL",
		},
		UPDATINTG: DatabaseStatusEnum{
			value: "UPDATINTG",
		},
		ABNORMAL: DatabaseStatusEnum{
			value: "ABNORMAL",
		},
		FREEZE: DatabaseStatusEnum{
			value: "FREEZE",
		},
		DELETING: DatabaseStatusEnum{
			value: "DELETING",
		},
	}
}

func (c DatabaseStatusEnum) Value() string {
	return c.value
}

func (c DatabaseStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseStatusEnum) UnmarshalJSON(b []byte) error {
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
