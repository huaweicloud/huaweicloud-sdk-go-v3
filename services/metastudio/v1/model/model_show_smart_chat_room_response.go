package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSmartChatRoomResponse Response Object
type ShowSmartChatRoomResponse struct {

	// 对话名称
	RoomName string `json:"room_name"`

	// 对话描述。
	RoomDescription *string `json:"room_description,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 数字人模型资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	// 机器人ID。获取方法请参考[创建应用](CreateRobot.xml)。
	RobotId *string `json:"robot_id,omitempty"`

	// **参数解释**： 并发路数。
	Concurrency *int32 `json:"concurrency,omitempty"`

	BackgroundConfig *BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	ChatSubtitleConfig *ChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	// 智能交互对话端配置。 * COMPUTER: 电脑端 * MOBILE: 手机端 * HUB: 大屏
	ChatVideoType *ShowSmartChatRoomResponseChatVideoType `json:"chat_video_type,omitempty"`

	// 对话ID。
	RoomId *string `json:"room_id,omitempty"`

	// 智能交互对话创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 智能交互对话更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`

	// 对话封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSmartChatRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmartChatRoomResponse struct{}"
	}

	return strings.Join([]string{"ShowSmartChatRoomResponse", string(data)}, " ")
}

type ShowSmartChatRoomResponseChatVideoType struct {
	value string
}

type ShowSmartChatRoomResponseChatVideoTypeEnum struct {
	COMPUTER ShowSmartChatRoomResponseChatVideoType
	MOBILE   ShowSmartChatRoomResponseChatVideoType
	HUB      ShowSmartChatRoomResponseChatVideoType
}

func GetShowSmartChatRoomResponseChatVideoTypeEnum() ShowSmartChatRoomResponseChatVideoTypeEnum {
	return ShowSmartChatRoomResponseChatVideoTypeEnum{
		COMPUTER: ShowSmartChatRoomResponseChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: ShowSmartChatRoomResponseChatVideoType{
			value: "MOBILE",
		},
		HUB: ShowSmartChatRoomResponseChatVideoType{
			value: "HUB",
		},
	}
}

func (c ShowSmartChatRoomResponseChatVideoType) Value() string {
	return c.value
}

func (c ShowSmartChatRoomResponseChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSmartChatRoomResponseChatVideoType) UnmarshalJSON(b []byte) error {
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
