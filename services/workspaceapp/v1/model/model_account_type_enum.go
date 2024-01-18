package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountTypeEnum 用户类型： * `USER` - 用户 * `USER_GROUP` - 用户组
type AccountTypeEnum struct {
	value string
}

type AccountTypeEnumEnum struct {
	USER       AccountTypeEnum
	USER_GROUP AccountTypeEnum
}

func GetAccountTypeEnumEnum() AccountTypeEnumEnum {
	return AccountTypeEnumEnum{
		USER: AccountTypeEnum{
			value: "USER",
		},
		USER_GROUP: AccountTypeEnum{
			value: "USER_GROUP",
		},
	}
}

func (c AccountTypeEnum) Value() string {
	return c.value
}

func (c AccountTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountTypeEnum) UnmarshalJSON(b []byte) error {
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
