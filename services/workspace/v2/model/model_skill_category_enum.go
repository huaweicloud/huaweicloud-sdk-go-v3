package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SkillCategoryEnum 技能分类枚举。
type SkillCategoryEnum struct {
	value string
}

type SkillCategoryEnumEnum struct {
	SYSTEM_CORE         SkillCategoryEnum
	AI_COGNITIVE        SkillCategoryEnum
	WEB_SEARCH          SkillCategoryEnum
	BROWSER_AUTOMATION  SkillCategoryEnum
	DEV_CODE            SkillCategoryEnum
	DOC_KNOWLEDGE       SkillCategoryEnum
	OFFICE_COLLAB       SkillCategoryEnum
	WORKFLOW_AUTOMATION SkillCategoryEnum
	MULTIMEDIA_CREATIVE SkillCategoryEnum
	SYSTEM_TOOLS        SkillCategoryEnum
}

func GetSkillCategoryEnumEnum() SkillCategoryEnumEnum {
	return SkillCategoryEnumEnum{
		SYSTEM_CORE: SkillCategoryEnum{
			value: "SYSTEM_CORE",
		},
		AI_COGNITIVE: SkillCategoryEnum{
			value: "AI_COGNITIVE",
		},
		WEB_SEARCH: SkillCategoryEnum{
			value: "WEB_SEARCH",
		},
		BROWSER_AUTOMATION: SkillCategoryEnum{
			value: "BROWSER_AUTOMATION",
		},
		DEV_CODE: SkillCategoryEnum{
			value: "DEV_CODE",
		},
		DOC_KNOWLEDGE: SkillCategoryEnum{
			value: "DOC_KNOWLEDGE",
		},
		OFFICE_COLLAB: SkillCategoryEnum{
			value: "OFFICE_COLLAB",
		},
		WORKFLOW_AUTOMATION: SkillCategoryEnum{
			value: "WORKFLOW_AUTOMATION",
		},
		MULTIMEDIA_CREATIVE: SkillCategoryEnum{
			value: "MULTIMEDIA_CREATIVE",
		},
		SYSTEM_TOOLS: SkillCategoryEnum{
			value: "SYSTEM_TOOLS",
		},
	}
}

func (c SkillCategoryEnum) Value() string {
	return c.value
}

func (c SkillCategoryEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SkillCategoryEnum) UnmarshalJSON(b []byte) error {
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
