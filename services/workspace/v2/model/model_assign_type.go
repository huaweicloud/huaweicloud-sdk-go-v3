package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssignType * `ALL_USER` - 全部用户 * `ASSIGN_USER` - 授权指定用户
type AssignType struct {
	value string
}

type AssignTypeEnum struct {
	ALL_USER    AssignType
	ASSIGN_USER AssignType
}

func GetAssignTypeEnum() AssignTypeEnum {
	return AssignTypeEnum{
		ALL_USER: AssignType{
			value: "ALL_USER",
		},
		ASSIGN_USER: AssignType{
			value: "ASSIGN_USER",
		},
	}
}

func (c AssignType) Value() string {
	return c.value
}

func (c AssignType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssignType) UnmarshalJSON(b []byte) error {
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
