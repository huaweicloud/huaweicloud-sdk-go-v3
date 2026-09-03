package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SkillSourceEnum 技能来源枚举。
type SkillSourceEnum struct {
	value string
}

type SkillSourceEnumEnum struct {
	CUSTOM      SkillSourceEnum
	OFFICIAL    SkillSourceEnum
	MARKETPLACE SkillSourceEnum
}

func GetSkillSourceEnumEnum() SkillSourceEnumEnum {
	return SkillSourceEnumEnum{
		CUSTOM: SkillSourceEnum{
			value: "CUSTOM",
		},
		OFFICIAL: SkillSourceEnum{
			value: "OFFICIAL",
		},
		MARKETPLACE: SkillSourceEnum{
			value: "MARKETPLACE",
		},
	}
}

func (c SkillSourceEnum) Value() string {
	return c.value
}

func (c SkillSourceEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SkillSourceEnum) UnmarshalJSON(b []byte) error {
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
