package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAiguardTemplateRuleResponse Response Object
type CreateAiguardTemplateRuleResponse struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 检测类别。text: 文本检测, image: 图片检测，为null时代表文字
	DetectType *CreateAiguardTemplateRuleResponseDetectType `json:"detect_type,omitempty"`

	// 检测目标，detect_type=image时有效。ocr代表图片中提取的文本，content代表图片语义
	DetectTarget *CreateAiguardTemplateRuleResponseDetectTarget `json:"detect_target,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则状态，可选：enable(启用), disable(禁用)
	Status *CreateAiguardTemplateRuleResponseStatus `json:"status,omitempty"`

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
	NotSupportedLanguageAction *CreateAiguardTemplateRuleResponseNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 规则来源，system代表系统，custom代表用户创建
	Source *CreateAiguardTemplateRuleResponseSource `json:"source,omitempty"`

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

func (o CreateAiguardTemplateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiguardTemplateRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAiguardTemplateRuleResponse", string(data)}, " ")
}

type CreateAiguardTemplateRuleResponseDetectType struct {
	value string
}

type CreateAiguardTemplateRuleResponseDetectTypeEnum struct {
	TEXT  CreateAiguardTemplateRuleResponseDetectType
	IMAGE CreateAiguardTemplateRuleResponseDetectType
}

func GetCreateAiguardTemplateRuleResponseDetectTypeEnum() CreateAiguardTemplateRuleResponseDetectTypeEnum {
	return CreateAiguardTemplateRuleResponseDetectTypeEnum{
		TEXT: CreateAiguardTemplateRuleResponseDetectType{
			value: "text",
		},
		IMAGE: CreateAiguardTemplateRuleResponseDetectType{
			value: "image",
		},
	}
}

func (c CreateAiguardTemplateRuleResponseDetectType) Value() string {
	return c.value
}

func (c CreateAiguardTemplateRuleResponseDetectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiguardTemplateRuleResponseDetectType) UnmarshalJSON(b []byte) error {
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

type CreateAiguardTemplateRuleResponseDetectTarget struct {
	value string
}

type CreateAiguardTemplateRuleResponseDetectTargetEnum struct {
	OCR     CreateAiguardTemplateRuleResponseDetectTarget
	CONTENT CreateAiguardTemplateRuleResponseDetectTarget
}

func GetCreateAiguardTemplateRuleResponseDetectTargetEnum() CreateAiguardTemplateRuleResponseDetectTargetEnum {
	return CreateAiguardTemplateRuleResponseDetectTargetEnum{
		OCR: CreateAiguardTemplateRuleResponseDetectTarget{
			value: "ocr",
		},
		CONTENT: CreateAiguardTemplateRuleResponseDetectTarget{
			value: "content",
		},
	}
}

func (c CreateAiguardTemplateRuleResponseDetectTarget) Value() string {
	return c.value
}

func (c CreateAiguardTemplateRuleResponseDetectTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiguardTemplateRuleResponseDetectTarget) UnmarshalJSON(b []byte) error {
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

type CreateAiguardTemplateRuleResponseStatus struct {
	value string
}

type CreateAiguardTemplateRuleResponseStatusEnum struct {
	ENABLE  CreateAiguardTemplateRuleResponseStatus
	DISABLE CreateAiguardTemplateRuleResponseStatus
}

func GetCreateAiguardTemplateRuleResponseStatusEnum() CreateAiguardTemplateRuleResponseStatusEnum {
	return CreateAiguardTemplateRuleResponseStatusEnum{
		ENABLE: CreateAiguardTemplateRuleResponseStatus{
			value: "enable",
		},
		DISABLE: CreateAiguardTemplateRuleResponseStatus{
			value: "disable",
		},
	}
}

func (c CreateAiguardTemplateRuleResponseStatus) Value() string {
	return c.value
}

func (c CreateAiguardTemplateRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiguardTemplateRuleResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateAiguardTemplateRuleResponseNotSupportedLanguageAction struct {
	value string
}

type CreateAiguardTemplateRuleResponseNotSupportedLanguageActionEnum struct {
	PASS  CreateAiguardTemplateRuleResponseNotSupportedLanguageAction
	BLOCK CreateAiguardTemplateRuleResponseNotSupportedLanguageAction
}

func GetCreateAiguardTemplateRuleResponseNotSupportedLanguageActionEnum() CreateAiguardTemplateRuleResponseNotSupportedLanguageActionEnum {
	return CreateAiguardTemplateRuleResponseNotSupportedLanguageActionEnum{
		PASS: CreateAiguardTemplateRuleResponseNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: CreateAiguardTemplateRuleResponseNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c CreateAiguardTemplateRuleResponseNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c CreateAiguardTemplateRuleResponseNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiguardTemplateRuleResponseNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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

type CreateAiguardTemplateRuleResponseSource struct {
	value string
}

type CreateAiguardTemplateRuleResponseSourceEnum struct {
	SYSTEM CreateAiguardTemplateRuleResponseSource
	CUSTOM CreateAiguardTemplateRuleResponseSource
}

func GetCreateAiguardTemplateRuleResponseSourceEnum() CreateAiguardTemplateRuleResponseSourceEnum {
	return CreateAiguardTemplateRuleResponseSourceEnum{
		SYSTEM: CreateAiguardTemplateRuleResponseSource{
			value: "system",
		},
		CUSTOM: CreateAiguardTemplateRuleResponseSource{
			value: "custom",
		},
	}
}

func (c CreateAiguardTemplateRuleResponseSource) Value() string {
	return c.value
}

func (c CreateAiguardTemplateRuleResponseSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAiguardTemplateRuleResponseSource) UnmarshalJSON(b []byte) error {
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
