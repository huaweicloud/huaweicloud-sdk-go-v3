package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	RobotId string `json:"robot_id"`

	// 并发路数。
	Concurrency *int32 `json:"concurrency,omitempty"`

	BackgroundConfig *BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	ChatSubtitleConfig *ChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`
}

func (o CreateSmartChatRoomReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmartChatRoomReq struct{}"
	}

	return strings.Join([]string{"CreateSmartChatRoomReq", string(data)}, " ")
}
