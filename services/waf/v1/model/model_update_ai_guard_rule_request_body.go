package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAiGuardRuleRequestBody 更新规则请求
type UpdateAiGuardRuleRequestBody struct {

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则状态，可选：enable(启用), disable(禁用)
	Status *UpdateAiGuardRuleRequestBodyStatus `json:"status,omitempty"`

	// 检测选项配置列表
	DetectOpts *[]DetectOpt `json:"detect_opts,omitempty"`

	Action *AiGuardRuleAction `json:"action,omitempty"`

	DetectLabels *DetectLabels `json:"detect_labels,omitempty"`

	// 不在支持语言范围的防护动作,pass: 放行, block: 拦截
	NotSupportedLanguageAction *UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 规则描述，最大长度1024
	Description *string `json:"description,omitempty"`

	// 自定义代答库ID（可选，当action.category=block时有效）
	CustomAnswersId *string `json:"custom_answers_id,omitempty"`

	// 自定义代答库名称（可选，当action.category=block时有效，与custom_answers_id对应）
	CustomAnswersName *string `json:"custom_answers_name,omitempty"`

	AbuseConfig *AbuseConfig `json:"abuse_config,omitempty"`

	TopicConfig *TopicConfig `json:"topic_config,omitempty"`
}

func (o UpdateAiGuardRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardRuleRequestBody", string(data)}, " ")
}

type UpdateAiGuardRuleRequestBodyStatus struct {
	value string
}

type UpdateAiGuardRuleRequestBodyStatusEnum struct {
	ENABLE  UpdateAiGuardRuleRequestBodyStatus
	DISABLE UpdateAiGuardRuleRequestBodyStatus
}

func GetUpdateAiGuardRuleRequestBodyStatusEnum() UpdateAiGuardRuleRequestBodyStatusEnum {
	return UpdateAiGuardRuleRequestBodyStatusEnum{
		ENABLE: UpdateAiGuardRuleRequestBodyStatus{
			value: "enable",
		},
		DISABLE: UpdateAiGuardRuleRequestBodyStatus{
			value: "disable",
		},
	}
}

func (c UpdateAiGuardRuleRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction struct {
	value string
}

type UpdateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum struct {
	PASS  UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction
	BLOCK UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction
}

func GetUpdateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum() UpdateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum {
	return UpdateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum{
		PASS: UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleRequestBodyNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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
