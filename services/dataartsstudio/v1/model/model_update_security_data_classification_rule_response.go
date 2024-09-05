package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSecurityDataClassificationRuleResponse Response Object
type UpdateSecurityDataClassificationRuleResponse struct {

	// 规则ID
	Uuid *string `json:"uuid,omitempty"`

	// 规则类型, CUSTOM, BUILTIN
	RuleType *UpdateSecurityDataClassificationRuleResponseRuleType `json:"rule_type,omitempty"`

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
	Method *UpdateSecurityDataClassificationRuleResponseMethod `json:"method,omitempty"`

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

func (o UpdateSecurityDataClassificationRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDataClassificationRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDataClassificationRuleResponse", string(data)}, " ")
}

type UpdateSecurityDataClassificationRuleResponseRuleType struct {
	value string
}

type UpdateSecurityDataClassificationRuleResponseRuleTypeEnum struct {
	CUSTOM  UpdateSecurityDataClassificationRuleResponseRuleType
	BUILTIN UpdateSecurityDataClassificationRuleResponseRuleType
}

func GetUpdateSecurityDataClassificationRuleResponseRuleTypeEnum() UpdateSecurityDataClassificationRuleResponseRuleTypeEnum {
	return UpdateSecurityDataClassificationRuleResponseRuleTypeEnum{
		CUSTOM: UpdateSecurityDataClassificationRuleResponseRuleType{
			value: "CUSTOM",
		},
		BUILTIN: UpdateSecurityDataClassificationRuleResponseRuleType{
			value: "BUILTIN",
		},
	}
}

func (c UpdateSecurityDataClassificationRuleResponseRuleType) Value() string {
	return c.value
}

func (c UpdateSecurityDataClassificationRuleResponseRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityDataClassificationRuleResponseRuleType) UnmarshalJSON(b []byte) error {
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

type UpdateSecurityDataClassificationRuleResponseMethod struct {
	value string
}

type UpdateSecurityDataClassificationRuleResponseMethodEnum struct {
	REGULAR UpdateSecurityDataClassificationRuleResponseMethod
	NONE    UpdateSecurityDataClassificationRuleResponseMethod
	DEFAULT UpdateSecurityDataClassificationRuleResponseMethod
	COMBINE UpdateSecurityDataClassificationRuleResponseMethod
}

func GetUpdateSecurityDataClassificationRuleResponseMethodEnum() UpdateSecurityDataClassificationRuleResponseMethodEnum {
	return UpdateSecurityDataClassificationRuleResponseMethodEnum{
		REGULAR: UpdateSecurityDataClassificationRuleResponseMethod{
			value: "REGULAR",
		},
		NONE: UpdateSecurityDataClassificationRuleResponseMethod{
			value: "NONE",
		},
		DEFAULT: UpdateSecurityDataClassificationRuleResponseMethod{
			value: "DEFAULT",
		},
		COMBINE: UpdateSecurityDataClassificationRuleResponseMethod{
			value: "COMBINE",
		},
	}
}

func (c UpdateSecurityDataClassificationRuleResponseMethod) Value() string {
	return c.value
}

func (c UpdateSecurityDataClassificationRuleResponseMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityDataClassificationRuleResponseMethod) UnmarshalJSON(b []byte) error {
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
