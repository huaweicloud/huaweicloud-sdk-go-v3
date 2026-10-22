package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAiguardTemplateRuleRequestBody 创建模板规则请求
type CreateAiguardTemplateRuleRequestBody struct {

	// 模板ID（8位字符串）
	TemplateId string `json:"template_id"`

	// 规则名称，非必选，通过模板生成
	Name *string `json:"name,omitempty"`

	Action *AiGuardRuleAction `json:"action,omitempty"`

	// 规则状态，非必选：enable(启用), disable(禁用)
	Status *CreateAiguardTemplateRuleRequestBodyStatus `json:"status,omitempty"`

	// 语言检测标签列表，只有使用语言模板添加时生效
	SupportedLanguageLabels *[]string `json:"supported_language_labels,omitempty"`

	// 其他语言检测动作，只有使用语言模板添加时生效。pass: 放行, block: 拦截
	NotSupportedLanguageAction *CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 自定义代答库ID（可选，当action.category=block时有效）
	CustomAnswersId *string `json:"custom_answers_id,omitempty"`

	// 自定义代答库名称（可选，当action.category=block时有效，与custom_answers_id对应）
	CustomAnswersName *string `json:"custom_answers_name,omitempty"`
}

func (o CreateAiguardTemplateRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiguardTemplateRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAiguardTemplateRuleRequestBody", string(data)}, " ")
}

type CreateAiguardTemplateRuleRequestBodyStatus struct {
	value string
}

type CreateAiguardTemplateRuleRequestBodyStatusEnum struct {
	ENABLE  CreateAiguardTemplateRuleRequestBodyStatus
	DISABLE CreateAiguardTemplateRuleRequestBodyStatus
}

func GetCreateAiguardTemplateRuleRequestBodyStatusEnum() CreateAiguardTemplateRuleRequestBodyStatusEnum {
	return CreateAiguardTemplateRuleRequestBodyStatusEnum{
		ENABLE: CreateAiguardTemplateRuleRequestBodyStatus{
			value: "enable",
		},
		DISABLE: CreateAiguardTemplateRuleRequestBodyStatus{
			value: "disable",
		},
	}
}

func (c CreateAiguardTemplateRuleRequestBodyStatus) Value() string {
	return c.value
}

func (c CreateAiguardTemplateRuleRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiguardTemplateRuleRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction struct {
	value string
}

type CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageActionEnum struct {
	PASS  CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction
	BLOCK CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction
}

func GetCreateAiguardTemplateRuleRequestBodyNotSupportedLanguageActionEnum() CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageActionEnum {
	return CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageActionEnum{
		PASS: CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiguardTemplateRuleRequestBodyNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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
