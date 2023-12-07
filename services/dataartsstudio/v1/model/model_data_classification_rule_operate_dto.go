package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataClassificationRuleOperateDto struct {

	// 规则类型, CUSTOM, BUILTIN
	RuleType *DataClassificationRuleOperateDtoRuleType `json:"rule_type,omitempty"`

	// 密级ID
	SecrecyLevelId *string `json:"secrecy_level_id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则方式, REGULAR, NONE, DEFAULT
	Method *DataClassificationRuleOperateDtoMethod `json:"method,omitempty"`

	// 内容表达式
	ContentExpression *string `json:"content_expression,omitempty"`

	// 列表达式
	ColumnExpression *string `json:"column_expression,omitempty"`

	// 备注表达式
	CommitExpression *string `json:"commit_expression,omitempty"`

	// 内置规则id
	BuiltinRuleId *string `json:"builtin_rule_id,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 分类ID
	CategoryId *string `json:"category_id,omitempty"`
}

func (o DataClassificationRuleOperateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationRuleOperateDto struct{}"
	}

	return strings.Join([]string{"DataClassificationRuleOperateDto", string(data)}, " ")
}

type DataClassificationRuleOperateDtoRuleType struct {
	value string
}

type DataClassificationRuleOperateDtoRuleTypeEnum struct {
	CUSTOM  DataClassificationRuleOperateDtoRuleType
	BUILTIN DataClassificationRuleOperateDtoRuleType
}

func GetDataClassificationRuleOperateDtoRuleTypeEnum() DataClassificationRuleOperateDtoRuleTypeEnum {
	return DataClassificationRuleOperateDtoRuleTypeEnum{
		CUSTOM: DataClassificationRuleOperateDtoRuleType{
			value: "CUSTOM",
		},
		BUILTIN: DataClassificationRuleOperateDtoRuleType{
			value: "BUILTIN",
		},
	}
}

func (c DataClassificationRuleOperateDtoRuleType) Value() string {
	return c.value
}

func (c DataClassificationRuleOperateDtoRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassificationRuleOperateDtoRuleType) UnmarshalJSON(b []byte) error {
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

type DataClassificationRuleOperateDtoMethod struct {
	value string
}

type DataClassificationRuleOperateDtoMethodEnum struct {
	REGULAR DataClassificationRuleOperateDtoMethod
	NONE    DataClassificationRuleOperateDtoMethod
	DEFAULT DataClassificationRuleOperateDtoMethod
}

func GetDataClassificationRuleOperateDtoMethodEnum() DataClassificationRuleOperateDtoMethodEnum {
	return DataClassificationRuleOperateDtoMethodEnum{
		REGULAR: DataClassificationRuleOperateDtoMethod{
			value: "REGULAR",
		},
		NONE: DataClassificationRuleOperateDtoMethod{
			value: "NONE",
		},
		DEFAULT: DataClassificationRuleOperateDtoMethod{
			value: "DEFAULT",
		},
	}
}

func (c DataClassificationRuleOperateDtoMethod) Value() string {
	return c.value
}

func (c DataClassificationRuleOperateDtoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassificationRuleOperateDtoMethod) UnmarshalJSON(b []byte) error {
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
