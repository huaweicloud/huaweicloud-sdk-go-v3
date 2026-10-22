package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAiGuardRuleRequestBody 创建规则请求
type CreateAiGuardRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// 检测选项配置列表
	DetectOpts []DetectOpt `json:"detect_opts"`

	Action *AiGuardRuleAction `json:"action"`

	DetectLabels *DetectLabels `json:"detect_labels"`

	// 不支持语言的防护动作，支持pass和block
	NotSupportedLanguageAction *CreateAiGuardRuleRequestBodyNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 检测目标，当detect_type=image时生效，ocr代表图片中提取的文本，content代表图片语义
	DetectTarget *CreateAiGuardRuleRequestBodyDetectTarget `json:"detect_target,omitempty"`

	// 规则描述，最大长度1024
	Description *string `json:"description,omitempty"`

	// 自定义代答库ID（可选，当action.category=block时有效）
	CustomAnswersId *string `json:"custom_answers_id,omitempty"`

	// 自定义代答库名称（可选，当action.category=block时有效，与custom_answers_id对应）
	CustomAnswersName *string `json:"custom_answers_name,omitempty"`

	AbuseConfig *AbuseConfig `json:"abuse_config,omitempty"`

	TopicConfig *TopicConfig `json:"topic_config,omitempty"`
}

func (o CreateAiGuardRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiGuardRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAiGuardRuleRequestBody", string(data)}, " ")
}

type CreateAiGuardRuleRequestBodyNotSupportedLanguageAction struct {
	value string
}

type CreateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum struct {
	PASS  CreateAiGuardRuleRequestBodyNotSupportedLanguageAction
	BLOCK CreateAiGuardRuleRequestBodyNotSupportedLanguageAction
}

func GetCreateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum() CreateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum {
	return CreateAiGuardRuleRequestBodyNotSupportedLanguageActionEnum{
		PASS: CreateAiGuardRuleRequestBodyNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: CreateAiGuardRuleRequestBodyNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c CreateAiGuardRuleRequestBodyNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c CreateAiGuardRuleRequestBodyNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiGuardRuleRequestBodyNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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

type CreateAiGuardRuleRequestBodyDetectTarget struct {
	value string
}

type CreateAiGuardRuleRequestBodyDetectTargetEnum struct {
	OCR     CreateAiGuardRuleRequestBodyDetectTarget
	CONTENT CreateAiGuardRuleRequestBodyDetectTarget
}

func GetCreateAiGuardRuleRequestBodyDetectTargetEnum() CreateAiGuardRuleRequestBodyDetectTargetEnum {
	return CreateAiGuardRuleRequestBodyDetectTargetEnum{
		OCR: CreateAiGuardRuleRequestBodyDetectTarget{
			value: "ocr",
		},
		CONTENT: CreateAiGuardRuleRequestBodyDetectTarget{
			value: "content",
		},
	}
}

func (c CreateAiGuardRuleRequestBodyDetectTarget) Value() string {
	return c.value
}

func (c CreateAiGuardRuleRequestBodyDetectTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiGuardRuleRequestBodyDetectTarget) UnmarshalJSON(b []byte) error {
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
