package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmartImageLayerConfig 素材图片图层配置。
type SmartImageLayerConfig struct {

	// 图片文件的URL。
	ImageUrl string `json:"image_url"`

	// **参数解释**： 图片显示时长，单位s。  显示时长规则为，若携带reply_texts、reply_audios，则与播放语音内容时长保持一致。若未携带，则与匹配的关键词语音内容时长保持一致。
	DisplayDuration *int32 `json:"display_duration,omitempty"`
}

func (o SmartImageLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartImageLayerConfig struct{}"
	}

	return strings.Join([]string{"SmartImageLayerConfig", string(data)}, " ")
}
