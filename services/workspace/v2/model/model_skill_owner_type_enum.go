package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SkillOwnerTypeEnum 技能所有者类型枚举。
type SkillOwnerTypeEnum struct {
	value string
}

type SkillOwnerTypeEnumEnum struct {
	COMMON SkillOwnerTypeEnum
	TENANT SkillOwnerTypeEnum
}

func GetSkillOwnerTypeEnumEnum() SkillOwnerTypeEnumEnum {
	return SkillOwnerTypeEnumEnum{
		COMMON: SkillOwnerTypeEnum{
			value: "COMMON",
		},
		TENANT: SkillOwnerTypeEnum{
			value: "TENANT",
		},
	}
}

func (c SkillOwnerTypeEnum) Value() string {
	return c.value
}

func (c SkillOwnerTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SkillOwnerTypeEnum) UnmarshalJSON(b []byte) error {
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
