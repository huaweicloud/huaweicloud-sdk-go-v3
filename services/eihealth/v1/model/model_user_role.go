package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserRole struct {
	value string
}

type UserRoleEnum struct {
	ADMIN    UserRole
	OPERATOR UserRole
}

func GetUserRoleEnum() UserRoleEnum {
	return UserRoleEnum{
		ADMIN: UserRole{
			value: "ADMIN",
		},
		OPERATOR: UserRole{
			value: "OPERATOR",
		},
	}
}

func (c UserRole) Value() string {
	return c.value
}

func (c UserRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserRole) UnmarshalJSON(b []byte) error {
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
