package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AiGuardRuleItem 规则详情
type AiGuardRuleItem struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 检测类别。text: 文本检测, image: 图片检测，为null时代表文字
	DetectType *AiGuardRuleItemDetectType `json:"detect_type,omitempty"`

	// 检测目标，detect_type=image时有效。ocr代表图片中提取的文本，content代表图片语义
	DetectTarget *AiGuardRuleItemDetectTarget `json:"detect_target,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则状态，可选：enable(启用), disable(禁用)
	Status *AiGuardRuleItemStatus `json:"status,omitempty"`

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
	NotSupportedLanguageAction *AiGuardRuleItemNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 规则来源，system代表系统，custom代表用户创建
	Source *AiGuardRuleItemSource `json:"source,omitempty"`

	// 当source=system时，source_id为默认模板规则的Id
	SourceId *string `json:"source_id,omitempty"`

	// 规则描述，最大长度1024
	Description *string `json:"description,omitempty"`

	// 自定义代答库ID（可选，当action.category=block时有效）
	CustomAnswersId *string `json:"custom_answers_id,omitempty"`

	// 自定义代答库名称（可选，当action.category=block时有效，与custom_answers_id对应）
	CustomAnswersName *string `json:"custom_answers_name,omitempty"`

	AbuseConfig *AbuseConfig `json:"abuse_config,omitempty"`

	TopicConfig *TopicConfig `json:"topic_config,omitempty"`
}

func (o AiGuardRuleItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiGuardRuleItem struct{}"
	}

	return strings.Join([]string{"AiGuardRuleItem", string(data)}, " ")
}

type AiGuardRuleItemDetectType struct {
	value string
}

type AiGuardRuleItemDetectTypeEnum struct {
	TEXT  AiGuardRuleItemDetectType
	IMAGE AiGuardRuleItemDetectType
}

func GetAiGuardRuleItemDetectTypeEnum() AiGuardRuleItemDetectTypeEnum {
	return AiGuardRuleItemDetectTypeEnum{
		TEXT: AiGuardRuleItemDetectType{
			value: "text",
		},
		IMAGE: AiGuardRuleItemDetectType{
			value: "image",
		},
	}
}

func (c AiGuardRuleItemDetectType) Value() string {
	return c.value
}

func (c AiGuardRuleItemDetectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AiGuardRuleItemDetectType) UnmarshalJSON(b []byte) error {
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

type AiGuardRuleItemDetectTarget struct {
	value string
}

type AiGuardRuleItemDetectTargetEnum struct {
	OCR     AiGuardRuleItemDetectTarget
	CONTENT AiGuardRuleItemDetectTarget
}

func GetAiGuardRuleItemDetectTargetEnum() AiGuardRuleItemDetectTargetEnum {
	return AiGuardRuleItemDetectTargetEnum{
		OCR: AiGuardRuleItemDetectTarget{
			value: "ocr",
		},
		CONTENT: AiGuardRuleItemDetectTarget{
			value: "content",
		},
	}
}

func (c AiGuardRuleItemDetectTarget) Value() string {
	return c.value
}

func (c AiGuardRuleItemDetectTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AiGuardRuleItemDetectTarget) UnmarshalJSON(b []byte) error {
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

type AiGuardRuleItemStatus struct {
	value string
}

type AiGuardRuleItemStatusEnum struct {
	ENABLE  AiGuardRuleItemStatus
	DISABLE AiGuardRuleItemStatus
}

func GetAiGuardRuleItemStatusEnum() AiGuardRuleItemStatusEnum {
	return AiGuardRuleItemStatusEnum{
		ENABLE: AiGuardRuleItemStatus{
			value: "enable",
		},
		DISABLE: AiGuardRuleItemStatus{
			value: "disable",
		},
	}
}

func (c AiGuardRuleItemStatus) Value() string {
	return c.value
}

func (c AiGuardRuleItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AiGuardRuleItemStatus) UnmarshalJSON(b []byte) error {
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

type AiGuardRuleItemNotSupportedLanguageAction struct {
	value string
}

type AiGuardRuleItemNotSupportedLanguageActionEnum struct {
	PASS  AiGuardRuleItemNotSupportedLanguageAction
	BLOCK AiGuardRuleItemNotSupportedLanguageAction
}

func GetAiGuardRuleItemNotSupportedLanguageActionEnum() AiGuardRuleItemNotSupportedLanguageActionEnum {
	return AiGuardRuleItemNotSupportedLanguageActionEnum{
		PASS: AiGuardRuleItemNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: AiGuardRuleItemNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c AiGuardRuleItemNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c AiGuardRuleItemNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AiGuardRuleItemNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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

type AiGuardRuleItemSource struct {
	value string
}

type AiGuardRuleItemSourceEnum struct {
	SYSTEM AiGuardRuleItemSource
	CUSTOM AiGuardRuleItemSource
}

func GetAiGuardRuleItemSourceEnum() AiGuardRuleItemSourceEnum {
	return AiGuardRuleItemSourceEnum{
		SYSTEM: AiGuardRuleItemSource{
			value: "system",
		},
		CUSTOM: AiGuardRuleItemSource{
			value: "custom",
		},
	}
}

func (c AiGuardRuleItemSource) Value() string {
	return c.value
}

func (c AiGuardRuleItemSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AiGuardRuleItemSource) UnmarshalJSON(b []byte) error {
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
