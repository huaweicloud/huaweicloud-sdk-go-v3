package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAiGuardRuleResponse Response Object
type UpdateAiGuardRuleResponse struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 检测类别。text: 文本检测, image: 图片检测，为null时代表文字
	DetectType *UpdateAiGuardRuleResponseDetectType `json:"detect_type,omitempty"`

	// 检测目标，detect_type=image时有效。ocr代表图片中提取的文本，content代表图片语义
	DetectTarget *UpdateAiGuardRuleResponseDetectTarget `json:"detect_target,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则状态，可选：enable(启用), disable(禁用)
	Status *UpdateAiGuardRuleResponseStatus `json:"status,omitempty"`

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
	NotSupportedLanguageAction *UpdateAiGuardRuleResponseNotSupportedLanguageAction `json:"not_supported_language_action,omitempty"`

	// 规则来源，system代表系统，custom代表用户创建
	Source *UpdateAiGuardRuleResponseSource `json:"source,omitempty"`

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

func (o UpdateAiGuardRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardRuleResponse", string(data)}, " ")
}

type UpdateAiGuardRuleResponseDetectType struct {
	value string
}

type UpdateAiGuardRuleResponseDetectTypeEnum struct {
	TEXT  UpdateAiGuardRuleResponseDetectType
	IMAGE UpdateAiGuardRuleResponseDetectType
}

func GetUpdateAiGuardRuleResponseDetectTypeEnum() UpdateAiGuardRuleResponseDetectTypeEnum {
	return UpdateAiGuardRuleResponseDetectTypeEnum{
		TEXT: UpdateAiGuardRuleResponseDetectType{
			value: "text",
		},
		IMAGE: UpdateAiGuardRuleResponseDetectType{
			value: "image",
		},
	}
}

func (c UpdateAiGuardRuleResponseDetectType) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleResponseDetectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleResponseDetectType) UnmarshalJSON(b []byte) error {
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

type UpdateAiGuardRuleResponseDetectTarget struct {
	value string
}

type UpdateAiGuardRuleResponseDetectTargetEnum struct {
	OCR     UpdateAiGuardRuleResponseDetectTarget
	CONTENT UpdateAiGuardRuleResponseDetectTarget
}

func GetUpdateAiGuardRuleResponseDetectTargetEnum() UpdateAiGuardRuleResponseDetectTargetEnum {
	return UpdateAiGuardRuleResponseDetectTargetEnum{
		OCR: UpdateAiGuardRuleResponseDetectTarget{
			value: "ocr",
		},
		CONTENT: UpdateAiGuardRuleResponseDetectTarget{
			value: "content",
		},
	}
}

func (c UpdateAiGuardRuleResponseDetectTarget) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleResponseDetectTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleResponseDetectTarget) UnmarshalJSON(b []byte) error {
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

type UpdateAiGuardRuleResponseStatus struct {
	value string
}

type UpdateAiGuardRuleResponseStatusEnum struct {
	ENABLE  UpdateAiGuardRuleResponseStatus
	DISABLE UpdateAiGuardRuleResponseStatus
}

func GetUpdateAiGuardRuleResponseStatusEnum() UpdateAiGuardRuleResponseStatusEnum {
	return UpdateAiGuardRuleResponseStatusEnum{
		ENABLE: UpdateAiGuardRuleResponseStatus{
			value: "enable",
		},
		DISABLE: UpdateAiGuardRuleResponseStatus{
			value: "disable",
		},
	}
}

func (c UpdateAiGuardRuleResponseStatus) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateAiGuardRuleResponseNotSupportedLanguageAction struct {
	value string
}

type UpdateAiGuardRuleResponseNotSupportedLanguageActionEnum struct {
	PASS  UpdateAiGuardRuleResponseNotSupportedLanguageAction
	BLOCK UpdateAiGuardRuleResponseNotSupportedLanguageAction
}

func GetUpdateAiGuardRuleResponseNotSupportedLanguageActionEnum() UpdateAiGuardRuleResponseNotSupportedLanguageActionEnum {
	return UpdateAiGuardRuleResponseNotSupportedLanguageActionEnum{
		PASS: UpdateAiGuardRuleResponseNotSupportedLanguageAction{
			value: "pass",
		},
		BLOCK: UpdateAiGuardRuleResponseNotSupportedLanguageAction{
			value: "block",
		},
	}
}

func (c UpdateAiGuardRuleResponseNotSupportedLanguageAction) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleResponseNotSupportedLanguageAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleResponseNotSupportedLanguageAction) UnmarshalJSON(b []byte) error {
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

type UpdateAiGuardRuleResponseSource struct {
	value string
}

type UpdateAiGuardRuleResponseSourceEnum struct {
	SYSTEM UpdateAiGuardRuleResponseSource
	CUSTOM UpdateAiGuardRuleResponseSource
}

func GetUpdateAiGuardRuleResponseSourceEnum() UpdateAiGuardRuleResponseSourceEnum {
	return UpdateAiGuardRuleResponseSourceEnum{
		SYSTEM: UpdateAiGuardRuleResponseSource{
			value: "system",
		},
		CUSTOM: UpdateAiGuardRuleResponseSource{
			value: "custom",
		},
	}
}

func (c UpdateAiGuardRuleResponseSource) Value() string {
	return c.value
}

func (c UpdateAiGuardRuleResponseSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAiGuardRuleResponseSource) UnmarshalJSON(b []byte) error {
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
