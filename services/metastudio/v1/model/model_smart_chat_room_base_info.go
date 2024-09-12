package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// **参数解释**： 并发路数。
	Concurrency *int32 `json:"concurrency,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o SmartChatRoomBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartChatRoomBaseInfo struct{}"
	}

	return strings.Join([]string{"SmartChatRoomBaseInfo", string(data)}, " ")
}
