package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthServerAccessMode 辅助认证服务接入模式 INTERNET：互联网接入 DEDICATED：专线接入 SYSTEM_DEFAULT：系统默认
type AuthServerAccessMode struct {
	value string
}

type AuthServerAccessModeEnum struct {
	INTERNET       AuthServerAccessMode
	DEDICATED      AuthServerAccessMode
	SYSTEM_DEFAULT AuthServerAccessMode
}

func GetAuthServerAccessModeEnum() AuthServerAccessModeEnum {
	return AuthServerAccessModeEnum{
		INTERNET: AuthServerAccessMode{
			value: "INTERNET",
		},
		DEDICATED: AuthServerAccessMode{
			value: "DEDICATED",
		},
		SYSTEM_DEFAULT: AuthServerAccessMode{
			value: "SYSTEM_DEFAULT",
		},
	}
}

func (c AuthServerAccessMode) Value() string {
	return c.value
}

func (c AuthServerAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthServerAccessMode) UnmarshalJSON(b []byte) error {
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
