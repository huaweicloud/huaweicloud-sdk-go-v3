package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSmartChatRoomRequestBody struct {

	// 对话名称
	RoomName string `json:"room_name"`

	// 对话描述。
	RoomDescription *string `json:"room_description,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 数字人模型资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	// 语音配置参数列表。
	VoiceConfigList *[]ChatVoiceConfig `json:"voice_config_list,omitempty"`

	// 机器人ID。获取方法请参考[创建应用](CreateRobot.xml)。
	RobotId *string `json:"robot_id,omitempty"`

	// **参数解释**： 并发路数。 **约束限制**： 默认没有并发路数，如果不配置并发数量，则无法启动智能交互对话任务。
	Concurrency *int32 `json:"concurrency,omitempty"`

	// 默认语言，智能交互接口使用。默认值CN。 * CN：简体中文。 * EN：英语。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	DefaultLanguage *CreateSmartChatRoomRequestBodyDefaultLanguage `json:"default_language,omitempty"`

	BackgroundConfig *BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	ChatSubtitleConfig *ChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	// 智能交互对话端配置。 * COMPUTER: 电脑端 * MOBILE: 手机端 * HUB: 大屏
	ChatVideoType *CreateSmartChatRoomRequestBodyChatVideoType `json:"chat_video_type,omitempty"`

	// **参数解释**： 静默退出时长。
	ExitMuteThreshold *int32 `json:"exit_mute_threshold,omitempty"`
}

func (o CreateSmartChatRoomRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmartChatRoomRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSmartChatRoomRequestBody", string(data)}, " ")
}

type CreateSmartChatRoomRequestBodyDefaultLanguage struct {
	value string
}

type CreateSmartChatRoomRequestBodyDefaultLanguageEnum struct {
	CN     CreateSmartChatRoomRequestBodyDefaultLanguage
	EN     CreateSmartChatRoomRequestBodyDefaultLanguage
	ESP    CreateSmartChatRoomRequestBodyDefaultLanguage
	POR    CreateSmartChatRoomRequestBodyDefaultLanguage
	ARABIC CreateSmartChatRoomRequestBodyDefaultLanguage
	THAI   CreateSmartChatRoomRequestBodyDefaultLanguage
}

func GetCreateSmartChatRoomRequestBodyDefaultLanguageEnum() CreateSmartChatRoomRequestBodyDefaultLanguageEnum {
	return CreateSmartChatRoomRequestBodyDefaultLanguageEnum{
		CN: CreateSmartChatRoomRequestBodyDefaultLanguage{
			value: "CN",
		},
		EN: CreateSmartChatRoomRequestBodyDefaultLanguage{
			value: "EN",
		},
		ESP: CreateSmartChatRoomRequestBodyDefaultLanguage{
			value: "ESP",
		},
		POR: CreateSmartChatRoomRequestBodyDefaultLanguage{
			value: "por",
		},
		ARABIC: CreateSmartChatRoomRequestBodyDefaultLanguage{
			value: "Arabic",
		},
		THAI: CreateSmartChatRoomRequestBodyDefaultLanguage{
			value: "Thai",
		},
	}
}

func (c CreateSmartChatRoomRequestBodyDefaultLanguage) Value() string {
	return c.value
}

func (c CreateSmartChatRoomRequestBodyDefaultLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSmartChatRoomRequestBodyDefaultLanguage) UnmarshalJSON(b []byte) error {
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

type CreateSmartChatRoomRequestBodyChatVideoType struct {
	value string
}

type CreateSmartChatRoomRequestBodyChatVideoTypeEnum struct {
	COMPUTER CreateSmartChatRoomRequestBodyChatVideoType
	MOBILE   CreateSmartChatRoomRequestBodyChatVideoType
	HUB      CreateSmartChatRoomRequestBodyChatVideoType
}

func GetCreateSmartChatRoomRequestBodyChatVideoTypeEnum() CreateSmartChatRoomRequestBodyChatVideoTypeEnum {
	return CreateSmartChatRoomRequestBodyChatVideoTypeEnum{
		COMPUTER: CreateSmartChatRoomRequestBodyChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: CreateSmartChatRoomRequestBodyChatVideoType{
			value: "MOBILE",
		},
		HUB: CreateSmartChatRoomRequestBodyChatVideoType{
			value: "HUB",
		},
	}
}

func (c CreateSmartChatRoomRequestBodyChatVideoType) Value() string {
	return c.value
}

func (c CreateSmartChatRoomRequestBodyChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSmartChatRoomRequestBodyChatVideoType) UnmarshalJSON(b []byte) error {
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
