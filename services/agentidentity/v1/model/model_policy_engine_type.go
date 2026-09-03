package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyEngineType 策略集的类型。GATEWAY - 用于网关的策略集，可绑定到网关实例。TOKEN_VAULT - 用于 TokenVault 的策略集，只能绑定到 TokenVault 实例。
type PolicyEngineType struct {
	value string
}

type PolicyEngineTypeEnum struct {
	GATEWAY     PolicyEngineType
	TOKEN_VAULT PolicyEngineType
}

func GetPolicyEngineTypeEnum() PolicyEngineTypeEnum {
	return PolicyEngineTypeEnum{
		GATEWAY: PolicyEngineType{
			value: "GATEWAY",
		},
		TOKEN_VAULT: PolicyEngineType{
			value: "TOKEN_VAULT",
		},
	}
}

func (c PolicyEngineType) Value() string {
	return c.value
}

func (c PolicyEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyEngineType) UnmarshalJSON(b []byte) error {
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
