package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SkillStatusEnum 技能状态枚举。
type SkillStatusEnum struct {
	value string
}

type SkillStatusEnumEnum struct {
	ACTIVE   SkillStatusEnum
	DISABLED SkillStatusEnum
}

func GetSkillStatusEnumEnum() SkillStatusEnumEnum {
	return SkillStatusEnumEnum{
		ACTIVE: SkillStatusEnum{
			value: "ACTIVE",
		},
		DISABLED: SkillStatusEnum{
			value: "DISABLED",
		},
	}
}

func (c SkillStatusEnum) Value() string {
	return c.value
}

func (c SkillStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SkillStatusEnum) UnmarshalJSON(b []byte) error {
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
