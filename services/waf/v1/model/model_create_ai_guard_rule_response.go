package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAiGuardRuleResponse Response Object
type CreateAiGuardRuleResponse struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 检测类别。text: 文本检测, image: 图片检测，为null时代表文字
	DetectType *CreateAiGuardRuleResponseDetectType `json:"detect_type,omitempty"`

	// 检测目标，detect_type=image时有效。ocr代表图片中提取的文本，content代表图片语义
	DetectTarget *CreateAiGuardRuleResponseDetectTarget `json:"detect_target,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则状态，可选：enable(启用), disable(禁用)
	Status *CreateAiGuardRuleResponseStatus `json:"status,omitempty"`

	// 检测选项配置列表
	DetectOpts *[]DetectOpt `json:"detect_opts,omitempty"`

	Action *AiGuardRuleAction `json:"action,omitempty"`

	DetectLabels *DetectLabels `json:"detect_labels,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	DetectLabelsInfo *DetectLabelsInfo `json:"detect_labels_info,omitempty"`

	// 不支持语言的检测防护动作。pass: 放行, block: 拦截
	NotSupportedLanguageAction *CreateAiGuardRuleResponseNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 规则来源，system代表系统，custom代表用户创建
	Source *CreateAiGuardRuleResponseSource `json:"source,omitempty"`

	// 当source=system时，source_id为默认模板规则的Id
	SourceId *string `json:"source_id,omitempty"`

	// 规则描述，最大长度1024
	Description *string `json:"description,omitempty"`

	// 自定义代答库ID（可选，当action.category=block时有效）
	CustomAnswersId *string `json:"custom_answers_id,omitempty"`

	// 自定义代答库名称（可选，当action.category=block时有效，与custom_answers_id对应）
	CustomAnswersName *string `json:"custom_answers_name,omitempty"`

	AbuseConfig *AbuseConfig `json:"abuse_config,omitempty"`

	TopicConfig    *TopicConfig `json:"topic_config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateAiGuardRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiGuardRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAiGuardRuleResponse", string(data)}, " ")
}

type CreateAiGuardRuleResponseDetectType struct {
	value string
}

type CreateAiGuardRuleResponseDetectTypeEnum struct {
	TEXT  CreateAiGuardRuleResponseDetectType
	IMAGE CreateAiGuardRuleResponseDetectType
}

func GetCreateAiGuardRuleResponseDetectTypeEnum() CreateAiGuardRuleResponseDetectTypeEnum {
	return CreateAiGuardRuleResponseDetectTypeEnum{
		TEXT: CreateAiGuardRuleResponseDetectType{
			value: "text",
		},
		IMAGE: CreateAiGuardRuleResponseDetectType{
			value: "image",
		},
	}
}

func (c CreateAiGuardRuleResponseDetectType) Value() string {
	return c.value
}

func (c CreateAiGuardRuleResponseDetectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiGuardRuleResponseDetectType) UnmarshalJSON(b []byte) error {
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

type CreateAiGuardRuleResponseDetectTarget struct {
	value string
}

type CreateAiGuardRuleResponseDetectTargetEnum struct {
	OCR     CreateAiGuardRuleResponseDetectTarget
	CONTENT CreateAiGuardRuleResponseDetectTarget
}

func GetCreateAiGuardRuleResponseDetectTargetEnum() CreateAiGuardRuleResponseDetectTargetEnum {
	return CreateAiGuardRuleResponseDetectTargetEnum{
		OCR: CreateAiGuardRuleResponseDetectTarget{
			value: "ocr",
		},
		CONTENT: CreateAiGuardRuleResponseDetectTarget{
			value: "content",
		},
	}
}

func (c CreateAiGuardRuleResponseDetectTarget) Value() string {
	return c.value
}

func (c CreateAiGuardRuleResponseDetectTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiGuardRuleResponseDetectTarget) UnmarshalJSON(b []byte) error {
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

type CreateAiGuardRuleResponseStatus struct {
	value string
}

type CreateAiGuardRuleResponseStatusEnum struct {
	ENABLE  CreateAiGuardRuleResponseStatus
	DISABLE CreateAiGuardRuleResponseStatus
}

func GetCreateAiGuardRuleResponseStatusEnum() CreateAiGuardRuleResponseStatusEnum {
	return CreateAiGuardRuleResponseStatusEnum{
		ENABLE: CreateAiGuardRuleResponseStatus{
			value: "enable",
		},
		DISABLE: CreateAiGuardRuleResponseStatus{
			value: "disable",
		},
	}
}

func (c CreateAiGuardRuleResponseStatus) Value() string {
	return c.value
}

func (c CreateAiGuardRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiGuardRuleResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateAiGuardRuleResponseNotSupportedLanguageAction struct {
	value string
}

type CreateAiGuardRuleResponseNotSupportedLanguageActionEnum struct {
	PASS  CreateAiGuardRuleResponseNotSupportedLanguageAction
	BLOCK CreateAiGuardRuleResponseNotSupportedLanguageAction
}

func GetCreateAiGuardRuleResponseNotSupportedLanguageActionEnum() CreateAiGuardRuleResponseNotSupportedLanguageActionEnum {
	return CreateAiGuardRuleResponseNotSupportedLanguageActionEnum{
		PASS: CreateAiGuardRuleResponseNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: CreateAiGuardRuleResponseNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c CreateAiGuardRuleResponseNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c CreateAiGuardRuleResponseNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiGuardRuleResponseNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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

type CreateAiGuardRuleResponseSource struct {
	value string
}

type CreateAiGuardRuleResponseSourceEnum struct {
	SYSTEM CreateAiGuardRuleResponseSource
	CUSTOM CreateAiGuardRuleResponseSource
}

func GetCreateAiGuardRuleResponseSourceEnum() CreateAiGuardRuleResponseSourceEnum {
	return CreateAiGuardRuleResponseSourceEnum{
		SYSTEM: CreateAiGuardRuleResponseSource{
			value: "system",
		},
		CUSTOM: CreateAiGuardRuleResponseSource{
			value: "custom",
		},
	}
}

func (c CreateAiGuardRuleResponseSource) Value() string {
	return c.value
}

func (c CreateAiGuardRuleResponseSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiGuardRuleResponseSource) UnmarshalJSON(b []byte) error {
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
