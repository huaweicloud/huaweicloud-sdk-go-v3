package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScriptCategory **参数解释**: 脚本分类 - RETRIEVE 检索 - ANALYSIS 分析  **约束限制** 不涉及 **取值范围**: - RETRIEVE - ANALYSIS  **默认值** 不涉及
type ScriptCategory struct {
	value string
}

type ScriptCategoryEnum struct {
	RETRIEVE ScriptCategory
	ANALYSIS ScriptCategory
}

func GetScriptCategoryEnum() ScriptCategoryEnum {
	return ScriptCategoryEnum{
		RETRIEVE: ScriptCategory{
			value: "RETRIEVE",
		},
		ANALYSIS: ScriptCategory{
			value: "ANALYSIS",
		},
	}
}

func (c ScriptCategory) Value() string {
	return c.value
}

func (c ScriptCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptCategory) UnmarshalJSON(b []byte) error {
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
