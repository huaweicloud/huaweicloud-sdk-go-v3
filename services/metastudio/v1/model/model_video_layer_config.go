package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoLayerConfig 素材视频图层配置。
type VideoLayerConfig struct {

	// **参数解释**： 视频文件的URL。 **约束限制**： * 仅直播支持外部URL，其他业务通过资产库查询获取，不支持外部URL。 **取值范围**： 字符长度1-2048位。 **默认取值**： 不涉及。
	VideoUrl *string `json:"video_url,omitempty"`

	// **参数解释**： 视频封面文件的URL。 **约束限制**： * 仅直播支持外部URL，其他业务通过资产库查询获取，不支持外部URL。 **取值范围**： 字符长度1-2048位。 **默认取值**： 不涉及。
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`

	// **参数解释**： 循环播放视频次数。  特殊取值： * 0：表示不播放 * -1：表示持续循环播放  **约束限制**： 不涉及。
	LoopCount *int32 `json:"loop_count,omitempty"`

	// **参数解释**： 按照百分比，调整视频素材的音量，取值为0-100。  特殊取值为0，表示不开启声音（默认值）。  **约束限制**： 不涉及。
	VideoSound *int32 `json:"video_sound,omitempty"`

	// **参数解释**： 是否播放完整个视频，true表示播放完整个视频，false表示当场景文本/音频结束时，视频也同时不再播放。  特殊取值： 默认值为false  **约束限制**： 不涉及。
	IsPlayTheEntireVideo *bool `json:"is_play_the_entire_video,omitempty"`
}

func (o VideoLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoLayerConfig struct{}"
	}

	return strings.Join([]string{"VideoLayerConfig", string(data)}, " ")
}
