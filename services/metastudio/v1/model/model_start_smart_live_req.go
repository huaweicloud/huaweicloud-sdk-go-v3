package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSmartLiveReq 数字人直播任务请求。
type StartSmartLiveReq struct {
	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	PlayPolicy *PlayPolicy `json:"play_policy,omitempty"`

	// RTMP视频推流第三方直播平台地址。
	OutputUrls *[]string `json:"output_urls,omitempty"`

	// RTMP视频推流第三方直播平台流秘钥，与推流地址对应。
	StreamKeys *[]string `json:"stream_keys,omitempty"`

	// 互动回调URL，含鉴权信息。
	InteractionCallbackUrl *string `json:"interaction_callback_url,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`
}

func (o StartSmartLiveReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSmartLiveReq struct{}"
	}

	return strings.Join([]string{"StartSmartLiveReq", string(data)}, " ")
}
