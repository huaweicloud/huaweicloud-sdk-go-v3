package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSmartLiveReq 数字人直播任务请求。
type StartSmartLiveReq struct {
	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	PlayPolicy *PlayPolicy `json:"play_policy,omitempty"`

	// 视频推流第三方直播平台地址。
	OutputUrls *[]string `json:"output_urls,omitempty"`
}

func (o StartSmartLiveReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSmartLiveReq struct{}"
	}

	return strings.Join([]string{"StartSmartLiveReq", string(data)}, " ")
}
