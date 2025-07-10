package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmartChatJobBaseInfo 数字人智能交互对话任务信息。
type SmartChatJobBaseInfo struct {

	// 数字人智能交互对话任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 数字人智能交互对话任务的状态。 * WAITING: 等待 * PROCESSING: 处理中 * SUCCEED: 成功 * FAILED: 失败 * DELETING: 删除中
	State *SmartChatJobBaseInfoState `json:"state,omitempty"`

	// 数字人智能交互对话时长，单位秒。
	Duration *float32 `json:"duration,omitempty"`

	// 数字人智能交互对话任务开始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人智能交互对话任务结束时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 数字人智能交互对话任务创建时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 数字人智能交互对话任务最后更新时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	LastupdateTime *string `json:"lastupdate_time,omitempty"`

	ChatSubtitleConfig *SmartChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	VideoConfig *SmartChatVideoConfig `json:"video_config,omitempty"`

	// 语音配置参数列表。
	VoiceConfigList *[]SmartChatVoiceConfig `json:"voice_config_list,omitempty"`

	// 智能交互对话端配置。 * COMPUTER: 电脑端 * MOBILE: 手机端 * HUB: 大屏
	ChatVideoType *SmartChatJobBaseInfoChatVideoType `json:"chat_video_type,omitempty"`

	// 是否透明背景
	IsTransparent *bool `json:"is_transparent,omitempty"`

	// 默认语言，智能交互接口使用。默认值CN。 * CN：中文。 * EN：英文。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	DefaultLanguage *SmartChatJobBaseInfoDefaultLanguage `json:"default_language,omitempty"`

	// clientId
	ClientId *string `json:"client_id,omitempty"`

	// 是否是资源池模式
	IsPoolMode *bool `json:"is_pool_mode,omitempty"`

	// 对话结束原因 * NORMAL：正常结束 * MUTE_TIMEOUT：静音超时
	JobFinishReason *SmartChatJobBaseInfoJobFinishReason `json:"job_finish_reason,omitempty"`
}

func (o SmartChatJobBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartChatJobBaseInfo struct{}"
	}

	return strings.Join([]string{"SmartChatJobBaseInfo", string(data)}, " ")
}

type SmartChatJobBaseInfoState struct {
	value string
}

type SmartChatJobBaseInfoStateEnum struct {
	WAITING    SmartChatJobBaseInfoState
	PROCESSING SmartChatJobBaseInfoState
	SUCCEED    SmartChatJobBaseInfoState
	FAILED     SmartChatJobBaseInfoState
	CANCELED   SmartChatJobBaseInfoState
	HEARTBEAT  SmartChatJobBaseInfoState
	IDLE       SmartChatJobBaseInfoState
	DELETING   SmartChatJobBaseInfoState
}

func GetSmartChatJobBaseInfoStateEnum() SmartChatJobBaseInfoStateEnum {
	return SmartChatJobBaseInfoStateEnum{
		WAITING: SmartChatJobBaseInfoState{
			value: "WAITING",
		},
		PROCESSING: SmartChatJobBaseInfoState{
			value: "PROCESSING",
		},
		SUCCEED: SmartChatJobBaseInfoState{
			value: "SUCCEED",
		},
		FAILED: SmartChatJobBaseInfoState{
			value: "FAILED",
		},
		CANCELED: SmartChatJobBaseInfoState{
			value: "CANCELED",
		},
		HEARTBEAT: SmartChatJobBaseInfoState{
			value: "HEARTBEAT",
		},
		IDLE: SmartChatJobBaseInfoState{
			value: "IDLE",
		},
		DELETING: SmartChatJobBaseInfoState{
			value: "DELETING",
		},
	}
}

func (c SmartChatJobBaseInfoState) Value() string {
	return c.value
}

func (c SmartChatJobBaseInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartChatJobBaseInfoState) UnmarshalJSON(b []byte) error {
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

type SmartChatJobBaseInfoChatVideoType struct {
	value string
}

type SmartChatJobBaseInfoChatVideoTypeEnum struct {
	COMPUTER SmartChatJobBaseInfoChatVideoType
	MOBILE   SmartChatJobBaseInfoChatVideoType
	HUB      SmartChatJobBaseInfoChatVideoType
}

func GetSmartChatJobBaseInfoChatVideoTypeEnum() SmartChatJobBaseInfoChatVideoTypeEnum {
	return SmartChatJobBaseInfoChatVideoTypeEnum{
		COMPUTER: SmartChatJobBaseInfoChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: SmartChatJobBaseInfoChatVideoType{
			value: "MOBILE",
		},
		HUB: SmartChatJobBaseInfoChatVideoType{
			value: "HUB",
		},
	}
}

func (c SmartChatJobBaseInfoChatVideoType) Value() string {
	return c.value
}

func (c SmartChatJobBaseInfoChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartChatJobBaseInfoChatVideoType) UnmarshalJSON(b []byte) error {
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

type SmartChatJobBaseInfoDefaultLanguage struct {
	value string
}

type SmartChatJobBaseInfoDefaultLanguageEnum struct {
	CN     SmartChatJobBaseInfoDefaultLanguage
	EN     SmartChatJobBaseInfoDefaultLanguage
	ESP    SmartChatJobBaseInfoDefaultLanguage
	POR    SmartChatJobBaseInfoDefaultLanguage
	ARABIC SmartChatJobBaseInfoDefaultLanguage
	THAI   SmartChatJobBaseInfoDefaultLanguage
}

func GetSmartChatJobBaseInfoDefaultLanguageEnum() SmartChatJobBaseInfoDefaultLanguageEnum {
	return SmartChatJobBaseInfoDefaultLanguageEnum{
		CN: SmartChatJobBaseInfoDefaultLanguage{
			value: "CN",
		},
		EN: SmartChatJobBaseInfoDefaultLanguage{
			value: "EN",
		},
		ESP: SmartChatJobBaseInfoDefaultLanguage{
			value: "ESP",
		},
		POR: SmartChatJobBaseInfoDefaultLanguage{
			value: "por",
		},
		ARABIC: SmartChatJobBaseInfoDefaultLanguage{
			value: "Arabic",
		},
		THAI: SmartChatJobBaseInfoDefaultLanguage{
			value: "Thai",
		},
	}
}

func (c SmartChatJobBaseInfoDefaultLanguage) Value() string {
	return c.value
}

func (c SmartChatJobBaseInfoDefaultLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartChatJobBaseInfoDefaultLanguage) UnmarshalJSON(b []byte) error {
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

type SmartChatJobBaseInfoJobFinishReason struct {
	value string
}

type SmartChatJobBaseInfoJobFinishReasonEnum struct {
	NORMAL       SmartChatJobBaseInfoJobFinishReason
	MUTE_TIMEOUT SmartChatJobBaseInfoJobFinishReason
}

func GetSmartChatJobBaseInfoJobFinishReasonEnum() SmartChatJobBaseInfoJobFinishReasonEnum {
	return SmartChatJobBaseInfoJobFinishReasonEnum{
		NORMAL: SmartChatJobBaseInfoJobFinishReason{
			value: "NORMAL",
		},
		MUTE_TIMEOUT: SmartChatJobBaseInfoJobFinishReason{
			value: "MUTE_TIMEOUT",
		},
	}
}

func (c SmartChatJobBaseInfoJobFinishReason) Value() string {
	return c.value
}

func (c SmartChatJobBaseInfoJobFinishReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartChatJobBaseInfoJobFinishReason) UnmarshalJSON(b []byte) error {
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
