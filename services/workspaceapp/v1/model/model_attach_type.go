package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AttachType 关联对象类型 * `USER` -  用户 * `USER_GROUP` - 用户组
type AttachType struct {
	value string
}

type AttachTypeEnum struct {
	USER       AttachType
	USER_GROUP AttachType
}

func GetAttachTypeEnum() AttachTypeEnum {
	return AttachTypeEnum{
		USER: AttachType{
			value: "USER",
		},
		USER_GROUP: AttachType{
			value: "USER_GROUP",
		},
	}
}

func (c AttachType) Value() string {
	return c.value
}

func (c AttachType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachType) UnmarshalJSON(b []byte) error {
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
