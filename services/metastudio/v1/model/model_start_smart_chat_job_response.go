package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSmartChatJobResponse Response Object
type StartSmartChatJobResponse struct {

	// 智能交互对话任务ID。
	JobId *string `json:"job_id,omitempty"`

	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`

	ChatSubtitleConfig *ChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	VideoConfig *ChatVideoConfigRsp `json:"video_config,omitempty"`

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
