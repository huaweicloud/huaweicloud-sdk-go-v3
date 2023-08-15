package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClaimMode 存储声明的类型 * `USER` -  用户目录 * `SHARE` - 共享目录
type ClaimMode struct {
	value string
}

type ClaimModeEnum struct {
	USER  ClaimMode
	SHARE ClaimMode
}

func GetClaimModeEnum() ClaimModeEnum {
	return ClaimModeEnum{
		USER: ClaimMode{
			value: "USER",
		},
		SHARE: ClaimMode{
			value: "SHARE",
		},
	}
}

func (c ClaimMode) Value() string {
	return c.value
}

func (c ClaimMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClaimMode) UnmarshalJSON(b []byte) error {
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
