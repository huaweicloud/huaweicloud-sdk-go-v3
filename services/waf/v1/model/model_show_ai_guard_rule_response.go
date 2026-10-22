package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAiGuardRuleResponse Response Object
type ShowAiGuardRuleResponse struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 检测类别。text: 文本检测, image: 图片检测，为null时代表文字
	DetectType *ShowAiGuardRuleResponseDetectType `json:"detect_type,omitempty"`

	// 检测目标，detect_type=image时有效。ocr代表图片中提取的文本，content代表图片语义
	DetectTarget *ShowAiGuardRuleResponseDetectTarget `json:"detect_target,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则状态，可选：enable(启用), disable(禁用)
	Status *ShowAiGuardRuleResponseStatus `json:"status,omitempty"`

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
	NotSupportedLanguageAction *ShowAiGuardRuleResponseNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 规则来源，system代表系统，custom代表用户创建
	Source *ShowAiGuardRuleResponseSource `json:"source,omitempty"`

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

func (o ShowAiGuardRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiGuardRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowAiGuardRuleResponse", string(data)}, " ")
}

type ShowAiGuardRuleResponseDetectType struct {
	value string
}

type ShowAiGuardRuleResponseDetectTypeEnum struct {
	TEXT  ShowAiGuardRuleResponseDetectType
	IMAGE ShowAiGuardRuleResponseDetectType
}

func GetShowAiGuardRuleResponseDetectTypeEnum() ShowAiGuardRuleResponseDetectTypeEnum {
	return ShowAiGuardRuleResponseDetectTypeEnum{
		TEXT: ShowAiGuardRuleResponseDetectType{
			value: "text",
		},
		IMAGE: ShowAiGuardRuleResponseDetectType{
			value: "image",
		},
	}
}

func (c ShowAiGuardRuleResponseDetectType) Value() string {
	return c.value
}

func (c ShowAiGuardRuleResponseDetectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAiGuardRuleResponseDetectType) UnmarshalJSON(b []byte) error {
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

type ShowAiGuardRuleResponseDetectTarget struct {
	value string
}

type ShowAiGuardRuleResponseDetectTargetEnum struct {
	OCR     ShowAiGuardRuleResponseDetectTarget
	CONTENT ShowAiGuardRuleResponseDetectTarget
}

func GetShowAiGuardRuleResponseDetectTargetEnum() ShowAiGuardRuleResponseDetectTargetEnum {
	return ShowAiGuardRuleResponseDetectTargetEnum{
		OCR: ShowAiGuardRuleResponseDetectTarget{
			value: "ocr",
		},
		CONTENT: ShowAiGuardRuleResponseDetectTarget{
			value: "content",
		},
	}
}

func (c ShowAiGuardRuleResponseDetectTarget) Value() string {
	return c.value
}

func (c ShowAiGuardRuleResponseDetectTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAiGuardRuleResponseDetectTarget) UnmarshalJSON(b []byte) error {
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

type ShowAiGuardRuleResponseStatus struct {
	value string
}

type ShowAiGuardRuleResponseStatusEnum struct {
	ENABLE  ShowAiGuardRuleResponseStatus
	DISABLE ShowAiGuardRuleResponseStatus
}

func GetShowAiGuardRuleResponseStatusEnum() ShowAiGuardRuleResponseStatusEnum {
	return ShowAiGuardRuleResponseStatusEnum{
		ENABLE: ShowAiGuardRuleResponseStatus{
			value: "enable",
		},
		DISABLE: ShowAiGuardRuleResponseStatus{
			value: "disable",
		},
	}
}

func (c ShowAiGuardRuleResponseStatus) Value() string {
	return c.value
}

func (c ShowAiGuardRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAiGuardRuleResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowAiGuardRuleResponseNotSupportedLanguageAction struct {
	value string
}

type ShowAiGuardRuleResponseNotSupportedLanguageActionEnum struct {
	PASS  ShowAiGuardRuleResponseNotSupportedLanguageAction
	BLOCK ShowAiGuardRuleResponseNotSupportedLanguageAction
}

func GetShowAiGuardRuleResponseNotSupportedLanguageActionEnum() ShowAiGuardRuleResponseNotSupportedLanguageActionEnum {
	return ShowAiGuardRuleResponseNotSupportedLanguageActionEnum{
		PASS: ShowAiGuardRuleResponseNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: ShowAiGuardRuleResponseNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c ShowAiGuardRuleResponseNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c ShowAiGuardRuleResponseNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAiGuardRuleResponseNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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

type ShowAiGuardRuleResponseSource struct {
	value string
}

type ShowAiGuardRuleResponseSourceEnum struct {
	SYSTEM ShowAiGuardRuleResponseSource
	CUSTOM ShowAiGuardRuleResponseSource
}

func GetShowAiGuardRuleResponseSourceEnum() ShowAiGuardRuleResponseSourceEnum {
	return ShowAiGuardRuleResponseSourceEnum{
		SYSTEM: ShowAiGuardRuleResponseSource{
			value: "system",
		},
		CUSTOM: ShowAiGuardRuleResponseSource{
			value: "custom",
		},
	}
}

func (c ShowAiGuardRuleResponseSource) Value() string {
	return c.value
}

func (c ShowAiGuardRuleResponseSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAiGuardRuleResponseSource) UnmarshalJSON(b []byte) error {
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
