package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataClassificationSingleRuleDto struct {

	// 规则序号,大写字母
	RuleCode string `json:"rule_code"`

	// 算法类型, REGEX,REGEX_INSENSITIVE,GROOVY,LENGTH_EQ,LENGTH_GT,LENGTH_LT,BUILTIN
	AlgorithmType DataClassificationSingleRuleDtoAlgorithmType `json:"algorithm_type"`

	// 匹配类型, CONTENT,COLUMN,COMMIT,TABLE_NAME,TABLE_COMMENT,DATABASE_NAME
	MatchType DataClassificationSingleRuleDtoMatchType `json:"match_type"`

	// expression
	Expression string `json:"expression"`

	// 内置规则ID
	BuiltinRuleId *string `json:"builtin_rule_id,omitempty"`
}

func (o DataClassificationSingleRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationSingleRuleDto struct{}"
	}

	return strings.Join([]string{"DataClassificationSingleRuleDto", string(data)}, " ")
}

type DataClassificationSingleRuleDtoAlgorithmType struct {
	value string
}

type DataClassificationSingleRuleDtoAlgorithmTypeEnum struct {
	REGEX             DataClassificationSingleRuleDtoAlgorithmType
	REGEX_INSENSITIVE DataClassificationSingleRuleDtoAlgorithmType
	GROOVY            DataClassificationSingleRuleDtoAlgorithmType
	LENGTH_EQ         DataClassificationSingleRuleDtoAlgorithmType
	LENGTH_GT         DataClassificationSingleRuleDtoAlgorithmType
	LENGTH_LT         DataClassificationSingleRuleDtoAlgorithmType
	BUILTIN           DataClassificationSingleRuleDtoAlgorithmType
}

func GetDataClassificationSingleRuleDtoAlgorithmTypeEnum() DataClassificationSingleRuleDtoAlgorithmTypeEnum {
	return DataClassificationSingleRuleDtoAlgorithmTypeEnum{
		REGEX: DataClassificationSingleRuleDtoAlgorithmType{
			value: "REGEX",
		},
		REGEX_INSENSITIVE: DataClassificationSingleRuleDtoAlgorithmType{
			value: "REGEX_INSENSITIVE",
		},
		GROOVY: DataClassificationSingleRuleDtoAlgorithmType{
			value: "GROOVY",
		},
		LENGTH_EQ: DataClassificationSingleRuleDtoAlgorithmType{
			value: "LENGTH_EQ",
		},
		LENGTH_GT: DataClassificationSingleRuleDtoAlgorithmType{
			value: "LENGTH_GT",
		},
		LENGTH_LT: DataClassificationSingleRuleDtoAlgorithmType{
			value: "LENGTH_LT",
		},
		BUILTIN: DataClassificationSingleRuleDtoAlgorithmType{
			value: "BUILTIN",
		},
	}
}

func (c DataClassificationSingleRuleDtoAlgorithmType) Value() string {
	return c.value
}

func (c DataClassificationSingleRuleDtoAlgorithmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassificationSingleRuleDtoAlgorithmType) UnmarshalJSON(b []byte) error {
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

type DataClassificationSingleRuleDtoMatchType struct {
	value string
}

type DataClassificationSingleRuleDtoMatchTypeEnum struct {
	CONTENT       DataClassificationSingleRuleDtoMatchType
	COLUMN        DataClassificationSingleRuleDtoMatchType
	COMMIT        DataClassificationSingleRuleDtoMatchType
	TABLE_NAME    DataClassificationSingleRuleDtoMatchType
	TABLE_COMMENT DataClassificationSingleRuleDtoMatchType
	DATABASE_NAME DataClassificationSingleRuleDtoMatchType
}

func GetDataClassificationSingleRuleDtoMatchTypeEnum() DataClassificationSingleRuleDtoMatchTypeEnum {
	return DataClassificationSingleRuleDtoMatchTypeEnum{
		CONTENT: DataClassificationSingleRuleDtoMatchType{
			value: "CONTENT",
		},
		COLUMN: DataClassificationSingleRuleDtoMatchType{
			value: "COLUMN",
		},
		COMMIT: DataClassificationSingleRuleDtoMatchType{
			value: "COMMIT",
		},
		TABLE_NAME: DataClassificationSingleRuleDtoMatchType{
			value: "TABLE_NAME",
		},
		TABLE_COMMENT: DataClassificationSingleRuleDtoMatchType{
			value: "TABLE_COMMENT",
		},
		DATABASE_NAME: DataClassificationSingleRuleDtoMatchType{
			value: "DATABASE_NAME",
		},
	}
}

func (c DataClassificationSingleRuleDtoMatchType) Value() string {
	return c.value
}

func (c DataClassificationSingleRuleDtoMatchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassificationSingleRuleDtoMatchType) UnmarshalJSON(b []byte) error {
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
