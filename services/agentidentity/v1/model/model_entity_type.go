package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EntityType The type of entity that can be attached to a policy engine.
type EntityType struct {
	value string
}

type EntityTypeEnum struct {
	GATEWAY     EntityType
	TOKEN_VAULT EntityType
}

func GetEntityTypeEnum() EntityTypeEnum {
	return EntityTypeEnum{
		GATEWAY: EntityType{
			value: "GATEWAY",
		},
		TOKEN_VAULT: EntityType{
			value: "TOKEN_VAULT",
		},
	}
}

func (c EntityType) Value() string {
	return c.value
}

func (c EntityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EntityType) UnmarshalJSON(b []byte) error {
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
