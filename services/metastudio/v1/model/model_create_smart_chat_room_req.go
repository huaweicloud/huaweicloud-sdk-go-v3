package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSmartChatRoomReq 创建智能交互对话配置。
type CreateSmartChatRoomReq struct {

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
	ChatVideoType *CreateSmartChatRoomReqChatVideoType `json:"chat_video_type,omitempty"`
}

func (o CreateSmartChatRoomReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmartChatRoomReq struct{}"
	}

	return strings.Join([]string{"CreateSmartChatRoomReq", string(data)}, " ")
}

type CreateSmartChatRoomReqChatVideoType struct {
	value string
}

type CreateSmartChatRoomReqChatVideoTypeEnum struct {
	COMPUTER CreateSmartChatRoomReqChatVideoType
	MOBILE   CreateSmartChatRoomReqChatVideoType
	HUB      CreateSmartChatRoomReqChatVideoType
}

func GetCreateSmartChatRoomReqChatVideoTypeEnum() CreateSmartChatRoomReqChatVideoTypeEnum {
	return CreateSmartChatRoomReqChatVideoTypeEnum{
		COMPUTER: CreateSmartChatRoomReqChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: CreateSmartChatRoomReqChatVideoType{
			value: "MOBILE",
		},
		HUB: CreateSmartChatRoomReqChatVideoType{
			value: "HUB",
		},
	}
}

func (c CreateSmartChatRoomReqChatVideoType) Value() string {
	return c.value
}

func (c CreateSmartChatRoomReqChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSmartChatRoomReqChatVideoType) UnmarshalJSON(b []byte) error {
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
