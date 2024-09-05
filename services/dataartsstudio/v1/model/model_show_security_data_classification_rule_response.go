package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityDataClassificationRuleResponse Response Object
type ShowSecurityDataClassificationRuleResponse struct {

	// 规则ID
	Uuid *string `json:"uuid,omitempty"`

	// 规则类型, CUSTOM, BUILTIN
	RuleType *ShowSecurityDataClassificationRuleResponseRuleType `json:"rule_type,omitempty"`

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
	Method *ShowSecurityDataClassificationRuleResponseMethod `json:"method,omitempty"`

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
	MatchType      *string `json:"match_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecurityDataClassificationRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityDataClassificationRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityDataClassificationRuleResponse", string(data)}, " ")
}

type ShowSecurityDataClassificationRuleResponseRuleType struct {
	value string
}

type ShowSecurityDataClassificationRuleResponseRuleTypeEnum struct {
	CUSTOM  ShowSecurityDataClassificationRuleResponseRuleType
	BUILTIN ShowSecurityDataClassificationRuleResponseRuleType
}

func GetShowSecurityDataClassificationRuleResponseRuleTypeEnum() ShowSecurityDataClassificationRuleResponseRuleTypeEnum {
	return ShowSecurityDataClassificationRuleResponseRuleTypeEnum{
		CUSTOM: ShowSecurityDataClassificationRuleResponseRuleType{
			value: "CUSTOM",
		},
		BUILTIN: ShowSecurityDataClassificationRuleResponseRuleType{
			value: "BUILTIN",
		},
	}
}

func (c ShowSecurityDataClassificationRuleResponseRuleType) Value() string {
	return c.value
}

func (c ShowSecurityDataClassificationRuleResponseRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityDataClassificationRuleResponseRuleType) UnmarshalJSON(b []byte) error {
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

type ShowSecurityDataClassificationRuleResponseMethod struct {
	value string
}

type ShowSecurityDataClassificationRuleResponseMethodEnum struct {
	REGULAR ShowSecurityDataClassificationRuleResponseMethod
	NONE    ShowSecurityDataClassificationRuleResponseMethod
	DEFAULT ShowSecurityDataClassificationRuleResponseMethod
	COMBINE ShowSecurityDataClassificationRuleResponseMethod
}

func GetShowSecurityDataClassificationRuleResponseMethodEnum() ShowSecurityDataClassificationRuleResponseMethodEnum {
	return ShowSecurityDataClassificationRuleResponseMethodEnum{
		REGULAR: ShowSecurityDataClassificationRuleResponseMethod{
			value: "REGULAR",
		},
		NONE: ShowSecurityDataClassificationRuleResponseMethod{
			value: "NONE",
		},
		DEFAULT: ShowSecurityDataClassificationRuleResponseMethod{
			value: "DEFAULT",
		},
		COMBINE: ShowSecurityDataClassificationRuleResponseMethod{
			value: "COMBINE",
		},
	}
}

func (c ShowSecurityDataClassificationRuleResponseMethod) Value() string {
	return c.value
}

func (c ShowSecurityDataClassificationRuleResponseMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityDataClassificationRuleResponseMethod) UnmarshalJSON(b []byte) error {
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
