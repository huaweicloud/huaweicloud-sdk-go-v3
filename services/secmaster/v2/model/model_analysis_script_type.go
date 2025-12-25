package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AnalysisScriptType **参数解释**: 分析脚本类型 - SEC_MASTER_SQL 安全云脑SQL - RETRIEVE_SQL 检索SQL  **约束限制** 不涉及 **取值范围**: - SEC_MASTER_SQL - RETRIEVE_SQL  **默认值** 不涉及
type AnalysisScriptType struct {
	value string
}

type AnalysisScriptTypeEnum struct {
	SEC_MASTER_SQL AnalysisScriptType
	RETRIEVE_SQL   AnalysisScriptType
}

func GetAnalysisScriptTypeEnum() AnalysisScriptTypeEnum {
	return AnalysisScriptTypeEnum{
		SEC_MASTER_SQL: AnalysisScriptType{
			value: "SEC_MASTER_SQL",
		},
		RETRIEVE_SQL: AnalysisScriptType{
			value: "RETRIEVE_SQL",
		},
	}
}

func (c AnalysisScriptType) Value() string {
	return c.value
}

func (c AnalysisScriptType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnalysisScriptType) UnmarshalJSON(b []byte) error {
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
