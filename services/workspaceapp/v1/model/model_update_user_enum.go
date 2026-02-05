package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateUserEnum 更新用户类型。 * `ADD` - 添加。 * `DELETE` - 按照应用组授权。
type UpdateUserEnum struct {
	value string
}

type UpdateUserEnumEnum struct {
	ADD    UpdateUserEnum
	DELETE UpdateUserEnum
}

func GetUpdateUserEnumEnum() UpdateUserEnumEnum {
	return UpdateUserEnumEnum{
		ADD: UpdateUserEnum{
			value: "ADD",
		},
		DELETE: UpdateUserEnum{
			value: "DELETE",
		},
	}
}

func (c UpdateUserEnum) Value() string {
	return c.value
}

func (c UpdateUserEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserEnum) UnmarshalJSON(b []byte) error {
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
