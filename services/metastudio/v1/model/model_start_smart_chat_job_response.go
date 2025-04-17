package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartSmartChatJobResponse Response Object
type StartSmartChatJobResponse struct {

	// 智能交互对话任务ID。
	JobId *string `json:"job_id,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`

	ChatSubtitleConfig *SmartChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	VideoConfig *SmartChatVideoConfig `json:"video_config,omitempty"`

	// 语音配置参数列表。
	VoiceConfigList *[]SmartChatVoiceConfig `json:"voice_config_list,omitempty"`

	// 智能交互对话端配置。 * COMPUTER: 电脑端 * MOBILE: 手机端 * HUB: 大屏
	ChatVideoType *StartSmartChatJobResponseChatVideoType `json:"chat_video_type,omitempty"`

	// 算力所在region。 * cn-north-4: 北京4 * cn-southwest-2: 贵阳1
	Region *string `json:"region,omitempty"`

	// 智能交互接入地址。
	ChatAccessAddress *string `json:"chat_access_address,omitempty"`

	// 智能交互Rest接口接入地址。
	ChatAccessRestAddress *string `json:"chat_access_rest_address,omitempty"`

	// 是否透明背景
	IsTransparent *bool `json:"is_transparent,omitempty"`

	// 默认语言，智能交互接口使用。默认值CN。 * CN：中文。 * EN：英文。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	DefaultLanguage *StartSmartChatJobResponseDefaultLanguage `json:"default_language,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartSmartChatJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSmartChatJobResponse struct{}"
	}

	return strings.Join([]string{"StartSmartChatJobResponse", string(data)}, " ")
}

type StartSmartChatJobResponseChatVideoType struct {
	value string
}

type StartSmartChatJobResponseChatVideoTypeEnum struct {
	COMPUTER StartSmartChatJobResponseChatVideoType
	MOBILE   StartSmartChatJobResponseChatVideoType
	HUB      StartSmartChatJobResponseChatVideoType
}

func GetStartSmartChatJobResponseChatVideoTypeEnum() StartSmartChatJobResponseChatVideoTypeEnum {
	return StartSmartChatJobResponseChatVideoTypeEnum{
		COMPUTER: StartSmartChatJobResponseChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: StartSmartChatJobResponseChatVideoType{
			value: "MOBILE",
		},
		HUB: StartSmartChatJobResponseChatVideoType{
			value: "HUB",
		},
	}
}

func (c StartSmartChatJobResponseChatVideoType) Value() string {
	return c.value
}

func (c StartSmartChatJobResponseChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSmartChatJobResponseChatVideoType) UnmarshalJSON(b []byte) error {
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

type StartSmartChatJobResponseDefaultLanguage struct {
	value string
}

type StartSmartChatJobResponseDefaultLanguageEnum struct {
	CN     StartSmartChatJobResponseDefaultLanguage
	EN     StartSmartChatJobResponseDefaultLanguage
	ESP    StartSmartChatJobResponseDefaultLanguage
	POR    StartSmartChatJobResponseDefaultLanguage
	ARABIC StartSmartChatJobResponseDefaultLanguage
	THAI   StartSmartChatJobResponseDefaultLanguage
}

func GetStartSmartChatJobResponseDefaultLanguageEnum() StartSmartChatJobResponseDefaultLanguageEnum {
	return StartSmartChatJobResponseDefaultLanguageEnum{
		CN: StartSmartChatJobResponseDefaultLanguage{
			value: "CN",
		},
		EN: StartSmartChatJobResponseDefaultLanguage{
			value: "EN",
		},
		ESP: StartSmartChatJobResponseDefaultLanguage{
			value: "ESP",
		},
		POR: StartSmartChatJobResponseDefaultLanguage{
			value: "por",
		},
		ARABIC: StartSmartChatJobResponseDefaultLanguage{
			value: "Arabic",
		},
		THAI: StartSmartChatJobResponseDefaultLanguage{
			value: "Thai",
		},
	}
}

func (c StartSmartChatJobResponseDefaultLanguage) Value() string {
	return c.value
}

func (c StartSmartChatJobResponseDefaultLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSmartChatJobResponseDefaultLanguage) UnmarshalJSON(b []byte) error {
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
