package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmartVideoLayerConfig 素材视频图层配置。
type SmartVideoLayerConfig struct {

	// 视频文件的URL。
	VideoUrl string `json:"video_url"`

	// 视频封面文件的URL。
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`

	// **参数解释**： 图片显示时长，单位s。  显示时长规则为，若携带reply_texts、reply_audios，则与播放语音内容时长保持一致。若未携带，则与匹配的关键词语音内容时长保持一致。
	DisplayDuration *int32 `json:"display_duration,omitempty"`
}

func (o SmartVideoLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartVideoLayerConfig struct{}"
	}

	return strings.Join([]string{"SmartVideoLayerConfig", string(data)}, " ")
}
