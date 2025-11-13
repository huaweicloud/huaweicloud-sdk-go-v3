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

	// 计费模式，默认值CONCURRENCY * CONCURRENCY：并发计费 * CLIENT：按接入端计费 * CLIENT_TOKENS: 按接入端计费（TOKENS）
	BillingMode *CreateSmartChatRoomRequestBodyBillingMode `json:"billing_mode,omitempty"`

	// 是否允许使用未分配的并发数（端模式下不能复用），默认不使用。
	ReuseResource *bool `json:"reuse_resource,omitempty"`

	// **参数解释**： 并发路数。 **约束限制**： 默认没有并发路数，如果不配置并发数量，则无法启动智能交互对话任务。
	Concurrency *int32 `json:"concurrency,omitempty"`

	// **参数解释**： 允许接入终端端数量。
	ClientNums *int32 `json:"client_nums,omitempty"`

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

	// 是否优先级加载模型资产
	EnableSemanticAction *bool `json:"enable_semantic_action,omitempty"`
}

func (o CreateSmartChatRoomRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmartChatRoomRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSmartChatRoomRequestBody", string(data)}, " ")
}

type CreateSmartChatRoomRequestBodyBillingMode struct {
	value string
}

type CreateSmartChatRoomRequestBodyBillingModeEnum struct {
	CONCURRENCY   CreateSmartChatRoomRequestBodyBillingMode
	CLIENT        CreateSmartChatRoomRequestBodyBillingMode
	CLIENT_TOKENS CreateSmartChatRoomRequestBodyBillingMode
}

func GetCreateSmartChatRoomRequestBodyBillingModeEnum() CreateSmartChatRoomRequestBodyBillingModeEnum {
	return CreateSmartChatRoomRequestBodyBillingModeEnum{
		CONCURRENCY: CreateSmartChatRoomRequestBodyBillingMode{
			value: "CONCURRENCY",
		},
		CLIENT: CreateSmartChatRoomRequestBodyBillingMode{
			value: "CLIENT",
		},
		CLIENT_TOKENS: CreateSmartChatRoomRequestBodyBillingMode{
			value: "CLIENT_TOKENS",
		},
	}
}

func (c CreateSmartChatRoomRequestBodyBillingMode) Value() string {
	return c.value
}

func (c CreateSmartChatRoomRequestBodyBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSmartChatRoomRequestBodyBillingMode) UnmarshalJSON(b []byte) error {
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
