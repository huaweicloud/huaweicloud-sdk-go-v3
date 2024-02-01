package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveAudioConfig 直播音频配置
type LiveAudioConfig struct {

	// 插入音频资产的资产id，外部资产信息无需填写
	AssetId *string `json:"asset_id,omitempty"`

	// 音频URL。仅支持MP3格式，大小<100MB。输出会自动转化为单声道16KHZ采样。
	AudioUrl *string `json:"audio_url,omitempty"`

	// 音频对应的字幕文件URL。仅SRT格式，大小<1MB。
	SubtitleUrl *string `json:"subtitle_url,omitempty"`
}

func (o LiveAudioConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveAudioConfig struct{}"
	}

	return strings.Join([]string{"LiveAudioConfig", string(data)}, " ")
}
