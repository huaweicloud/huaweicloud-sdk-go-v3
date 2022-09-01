package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 在线会场信息。
type RealTimeParticipant struct {

	// 与会者标识。
	Pid *string `json:"pid,omitempty" xml:"pid"`

	// 与会者名称或昵称，长度限制为96个字符。
	Name *string `json:"name,omitempty" xml:"name"`

	// 与会者设备的注册号码（可支持SIP、TEL号码格式）。最大不超过127个字符。。
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// 用户状态。若会场未入会或已离会，则不会显示于在线会场列表。 - 0: 会议中。 - 1: 正在呼叫。 - 2: 正在加入会议。
	State *int32 `json:"state,omitempty" xml:"state"`

	// 音视频能力。 - 0: 音频。 - 1: 视频。
	Video *int32 `json:"video,omitempty" xml:"video"`

	// 麦克风状态。 - 0: 麦克风打开。 - 1: 麦克风关闭。
	Mute *int32 `json:"mute,omitempty" xml:"mute"`

	// 与会者举手状态。 - 0: 未举手。 - 1: 举手。
	Hand *int32 `json:"hand,omitempty" xml:"hand"`
}

func (o RealTimeParticipant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeParticipant struct{}"
	}

	return strings.Join([]string{"RealTimeParticipant", string(data)}, " ")
}
