package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoLayerConfig 素材视频图层配置。
type VideoLayerConfig struct {

	// **参数解释**： 视频文件的URL。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-2048位。 **默认取值**： 不涉及。
	VideoUrl *string `json:"video_url,omitempty"`

	// **参数解释**： 视频封面文件的URL。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-2048位。 **默认取值**： 不涉及。
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`

	// **参数解释**： 循环播放视频次数。  特殊取值： * 0：表示不播放 * -1：表示持续循环播放  **约束限制**： 不涉及。
	LoopCount *int32 `json:"loop_count,omitempty"`
}

func (o VideoLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoLayerConfig struct{}"
	}

	return strings.Join([]string{"VideoLayerConfig", string(data)}, " ")
}
