package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessKeyStatus 访问密钥状态，可以为“启用”（active）或“停用”（inactive）。
type AccessKeyStatus struct {
	value string
}

type AccessKeyStatusEnum struct {
	ACTIVE   AccessKeyStatus
	INACTIVE AccessKeyStatus
}

func GetAccessKeyStatusEnum() AccessKeyStatusEnum {
	return AccessKeyStatusEnum{
		ACTIVE: AccessKeyStatus{
			value: "active",
		},
		INACTIVE: AccessKeyStatus{
			value: "inactive",
		},
	}
}

func (c AccessKeyStatus) Value() string {
	return c.value
}

func (c AccessKeyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessKeyStatus) UnmarshalJSON(b []byte) error {
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
