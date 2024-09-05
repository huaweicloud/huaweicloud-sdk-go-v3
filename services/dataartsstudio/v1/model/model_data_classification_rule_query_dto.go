package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataClassificationRuleQueryDto struct {

	// 规则ID
	Uuid *string `json:"uuid,omitempty"`

	// 规则类型, CUSTOM, BUILTIN
	RuleType *DataClassificationRuleQueryDtoRuleType `json:"rule_type,omitempty"`

	// 密级
	SecrecyLevel *string `json:"secrecy_level,omitempty"`

	// 密级层级
	SecrecyLevelNum *int64 `json:"secrecy_level_num,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// guid
	Guid *string `json:"guid,omitempty"`

	// 规则是否开启
	Enable *bool `json:"enable,omitempty"`

	// 规则方式, REGULAR, NONE, DEFAULT, COMBINE
	Method *DataClassificationRuleQueryDtoMethod `json:"method,omitempty"`

	// 内容表达式
	ContentExpression *string `json:"content_expression,omitempty"`

	// 列表达式
	ColumnExpression *string `json:"column_expression,omitempty"`

	// 备注表达式
	CommitExpression *string `json:"commit_expression,omitempty"`

	// 条件表达式
	CombineExpression *string `json:"combine_expression,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 策略创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 策略创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 策略更新人
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 策略更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 内置规则ID
	BuiltinRuleId *string `json:"builtin_rule_id,omitempty"`

	// 分类ID
	CategoryId *string `json:"category_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 匹配类型
	MatchType *string `json:"match_type,omitempty"`
}

func (o DataClassificationRuleQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationRuleQueryDto struct{}"
	}

	return strings.Join([]string{"DataClassificationRuleQueryDto", string(data)}, " ")
}

type DataClassificationRuleQueryDtoRuleType struct {
	value string
}

type DataClassificationRuleQueryDtoRuleTypeEnum struct {
	CUSTOM  DataClassificationRuleQueryDtoRuleType
	BUILTIN DataClassificationRuleQueryDtoRuleType
}

func GetDataClassificationRuleQueryDtoRuleTypeEnum() DataClassificationRuleQueryDtoRuleTypeEnum {
	return DataClassificationRuleQueryDtoRuleTypeEnum{
		CUSTOM: DataClassificationRuleQueryDtoRuleType{
			value: "CUSTOM",
		},
		BUILTIN: DataClassificationRuleQueryDtoRuleType{
			value: "BUILTIN",
		},
	}
}

func (c DataClassificationRuleQueryDtoRuleType) Value() string {
	return c.value
}

func (c DataClassificationRuleQueryDtoRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassificationRuleQueryDtoRuleType) UnmarshalJSON(b []byte) error {
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

type DataClassificationRuleQueryDtoMethod struct {
	value string
}

type DataClassificationRuleQueryDtoMethodEnum struct {
	REGULAR DataClassificationRuleQueryDtoMethod
	NONE    DataClassificationRuleQueryDtoMethod
	DEFAULT DataClassificationRuleQueryDtoMethod
	COMBINE DataClassificationRuleQueryDtoMethod
}

func GetDataClassificationRuleQueryDtoMethodEnum() DataClassificationRuleQueryDtoMethodEnum {
	return DataClassificationRuleQueryDtoMethodEnum{
		REGULAR: DataClassificationRuleQueryDtoMethod{
			value: "REGULAR",
		},
		NONE: DataClassificationRuleQueryDtoMethod{
			value: "NONE",
		},
		DEFAULT: DataClassificationRuleQueryDtoMethod{
			value: "DEFAULT",
		},
		COMBINE: DataClassificationRuleQueryDtoMethod{
			value: "COMBINE",
		},
	}
}

func (c DataClassificationRuleQueryDtoMethod) Value() string {
	return c.value
}

func (c DataClassificationRuleQueryDtoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassificationRuleQueryDtoMethod) UnmarshalJSON(b []byte) error {
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
