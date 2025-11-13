package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSmartChatRoomResponse Response Object
type UpdateSmartChatRoomResponse struct {

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
	BillingMode *UpdateSmartChatRoomResponseBillingMode `json:"billing_mode,omitempty"`

	// 是否允许使用未分配的并发数（端模式下不能复用），默认不使用。
	ReuseResource *bool `json:"reuse_resource,omitempty"`

	// **参数解释**： 并发路数。 **约束限制**： 默认没有并发路数，如果不配置并发数量，则无法启动智能交互对话任务。
	Concurrency *int32 `json:"concurrency,omitempty"`

	// **参数解释**： 允许接入终端端数量。
	ClientNums *int32 `json:"client_nums,omitempty"`

	// 默认语言，智能交互接口使用。默认值CN。 * CN：简体中文。 * EN：英语。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	DefaultLanguage *UpdateSmartChatRoomResponseDefaultLanguage `json:"default_language,omitempty"`

	BackgroundConfig *BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	ChatSubtitleConfig *ChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	// 智能交互对话端配置。 * COMPUTER: 电脑端 * MOBILE: 手机端 * HUB: 大屏
	ChatVideoType *UpdateSmartChatRoomResponseChatVideoType `json:"chat_video_type,omitempty"`

	// **参数解释**： 静默退出时长。
	ExitMuteThreshold *int32 `json:"exit_mute_threshold,omitempty"`

	// 是否优先级加载模型资产
	EnableSemanticAction *bool `json:"enable_semantic_action,omitempty"`

	// 对话ID。
	RoomId *string `json:"room_id,omitempty"`

	// 智能交互对话创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 智能交互对话更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`

	// 对话封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	// 是否是资源池模式
	IsPoolMode *bool `json:"is_pool_mode,omitempty"`

	// 资源配置。
	ChatResourceConfig *[]ChatResourceConfigInfo `json:"chat_resource_config,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSmartChatRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSmartChatRoomResponse struct{}"
	}

	return strings.Join([]string{"UpdateSmartChatRoomResponse", string(data)}, " ")
}

type UpdateSmartChatRoomResponseBillingMode struct {
	value string
}

type UpdateSmartChatRoomResponseBillingModeEnum struct {
	CONCURRENCY   UpdateSmartChatRoomResponseBillingMode
	CLIENT        UpdateSmartChatRoomResponseBillingMode
	CLIENT_TOKENS UpdateSmartChatRoomResponseBillingMode
}

func GetUpdateSmartChatRoomResponseBillingModeEnum() UpdateSmartChatRoomResponseBillingModeEnum {
	return UpdateSmartChatRoomResponseBillingModeEnum{
		CONCURRENCY: UpdateSmartChatRoomResponseBillingMode{
			value: "CONCURRENCY",
		},
		CLIENT: UpdateSmartChatRoomResponseBillingMode{
			value: "CLIENT",
		},
		CLIENT_TOKENS: UpdateSmartChatRoomResponseBillingMode{
			value: "CLIENT_TOKENS",
		},
	}
}

func (c UpdateSmartChatRoomResponseBillingMode) Value() string {
	return c.value
}

func (c UpdateSmartChatRoomResponseBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSmartChatRoomResponseBillingMode) UnmarshalJSON(b []byte) error {
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

type UpdateSmartChatRoomResponseDefaultLanguage struct {
	value string
}

type UpdateSmartChatRoomResponseDefaultLanguageEnum struct {
	CN     UpdateSmartChatRoomResponseDefaultLanguage
	EN     UpdateSmartChatRoomResponseDefaultLanguage
	ESP    UpdateSmartChatRoomResponseDefaultLanguage
	POR    UpdateSmartChatRoomResponseDefaultLanguage
	ARABIC UpdateSmartChatRoomResponseDefaultLanguage
	THAI   UpdateSmartChatRoomResponseDefaultLanguage
}

func GetUpdateSmartChatRoomResponseDefaultLanguageEnum() UpdateSmartChatRoomResponseDefaultLanguageEnum {
	return UpdateSmartChatRoomResponseDefaultLanguageEnum{
		CN: UpdateSmartChatRoomResponseDefaultLanguage{
			value: "CN",
		},
		EN: UpdateSmartChatRoomResponseDefaultLanguage{
			value: "EN",
		},
		ESP: UpdateSmartChatRoomResponseDefaultLanguage{
			value: "ESP",
		},
		POR: UpdateSmartChatRoomResponseDefaultLanguage{
			value: "por",
		},
		ARABIC: UpdateSmartChatRoomResponseDefaultLanguage{
			value: "Arabic",
		},
		THAI: UpdateSmartChatRoomResponseDefaultLanguage{
			value: "Thai",
		},
	}
}

func (c UpdateSmartChatRoomResponseDefaultLanguage) Value() string {
	return c.value
}

func (c UpdateSmartChatRoomResponseDefaultLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSmartChatRoomResponseDefaultLanguage) UnmarshalJSON(b []byte) error {
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

type UpdateSmartChatRoomResponseChatVideoType struct {
	value string
}

type UpdateSmartChatRoomResponseChatVideoTypeEnum struct {
	COMPUTER UpdateSmartChatRoomResponseChatVideoType
	MOBILE   UpdateSmartChatRoomResponseChatVideoType
	HUB      UpdateSmartChatRoomResponseChatVideoType
}

func GetUpdateSmartChatRoomResponseChatVideoTypeEnum() UpdateSmartChatRoomResponseChatVideoTypeEnum {
	return UpdateSmartChatRoomResponseChatVideoTypeEnum{
		COMPUTER: UpdateSmartChatRoomResponseChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: UpdateSmartChatRoomResponseChatVideoType{
			value: "MOBILE",
		},
		HUB: UpdateSmartChatRoomResponseChatVideoType{
			value: "HUB",
		},
	}
}

func (c UpdateSmartChatRoomResponseChatVideoType) Value() string {
	return c.value
}

func (c UpdateSmartChatRoomResponseChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSmartChatRoomResponseChatVideoType) UnmarshalJSON(b []byte) error {
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
