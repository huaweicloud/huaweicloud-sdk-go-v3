package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SmartChatRoomBaseInfo struct {

	// 智能交互对话ID
	RoomId *string `json:"room_id,omitempty"`

	// 智能交互对话名称
	RoomName *string `json:"room_name,omitempty"`

	// 智能交互对话描述。
	RoomDescription *string `json:"room_description,omitempty"`

	// 机器人ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 对话封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	ModelInfos *ModelInfo `json:"model_infos,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	// 计费模式，默认值CONCURRENCY * CONCURRENCY：并发计费 * CLIENT：按接入端计费 * CLIENT_TOKENS: 按接入端计费（TOKENS）
	BillingMode *SmartChatRoomBaseInfoBillingMode `json:"billing_mode,omitempty"`

	// 是否允许使用未分配的并发数（端模式下不能复用），默认不使用。
	ReuseResource *bool `json:"reuse_resource,omitempty"`

	// **参数解释**： 并发路数。
	Concurrency *int32 `json:"concurrency,omitempty"`

	// **参数解释**： 允许接入终端端数量。
	ClientNums *int32 `json:"client_nums,omitempty"`

	// 语音配置参数列表。
	VoiceConfigList *[]VoiceConfigRsp `json:"voice_config_list,omitempty"`

	// 默认语言，智能交互接口使用。默认值CN。 * CN：简体中文。 * EN：英语。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	DefaultLanguage *SmartChatRoomBaseInfoDefaultLanguage `json:"default_language,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 静默退出时长。
	ExitMuteThreshold *int32 `json:"exit_mute_threshold,omitempty"`
}

func (o SmartChatRoomBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartChatRoomBaseInfo struct{}"
	}

	return strings.Join([]string{"SmartChatRoomBaseInfo", string(data)}, " ")
}

type SmartChatRoomBaseInfoBillingMode struct {
	value string
}

type SmartChatRoomBaseInfoBillingModeEnum struct {
	CONCURRENCY   SmartChatRoomBaseInfoBillingMode
	CLIENT        SmartChatRoomBaseInfoBillingMode
	CLIENT_TOKENS SmartChatRoomBaseInfoBillingMode
}

func GetSmartChatRoomBaseInfoBillingModeEnum() SmartChatRoomBaseInfoBillingModeEnum {
	return SmartChatRoomBaseInfoBillingModeEnum{
		CONCURRENCY: SmartChatRoomBaseInfoBillingMode{
			value: "CONCURRENCY",
		},
		CLIENT: SmartChatRoomBaseInfoBillingMode{
			value: "CLIENT",
		},
		CLIENT_TOKENS: SmartChatRoomBaseInfoBillingMode{
			value: "CLIENT_TOKENS",
		},
	}
}

func (c SmartChatRoomBaseInfoBillingMode) Value() string {
	return c.value
}

func (c SmartChatRoomBaseInfoBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartChatRoomBaseInfoBillingMode) UnmarshalJSON(b []byte) error {
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

type SmartChatRoomBaseInfoDefaultLanguage struct {
	value string
}

type SmartChatRoomBaseInfoDefaultLanguageEnum struct {
	CN     SmartChatRoomBaseInfoDefaultLanguage
	EN     SmartChatRoomBaseInfoDefaultLanguage
	ESP    SmartChatRoomBaseInfoDefaultLanguage
	POR    SmartChatRoomBaseInfoDefaultLanguage
	ARABIC SmartChatRoomBaseInfoDefaultLanguage
	THAI   SmartChatRoomBaseInfoDefaultLanguage
}

func GetSmartChatRoomBaseInfoDefaultLanguageEnum() SmartChatRoomBaseInfoDefaultLanguageEnum {
	return SmartChatRoomBaseInfoDefaultLanguageEnum{
		CN: SmartChatRoomBaseInfoDefaultLanguage{
			value: "CN",
		},
		EN: SmartChatRoomBaseInfoDefaultLanguage{
			value: "EN",
		},
		ESP: SmartChatRoomBaseInfoDefaultLanguage{
			value: "ESP",
		},
		POR: SmartChatRoomBaseInfoDefaultLanguage{
			value: "por",
		},
		ARABIC: SmartChatRoomBaseInfoDefaultLanguage{
			value: "Arabic",
		},
		THAI: SmartChatRoomBaseInfoDefaultLanguage{
			value: "Thai",
		},
	}
}

func (c SmartChatRoomBaseInfoDefaultLanguage) Value() string {
	return c.value
}

func (c SmartChatRoomBaseInfoDefaultLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartChatRoomBaseInfoDefaultLanguage) UnmarshalJSON(b []byte) error {
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
