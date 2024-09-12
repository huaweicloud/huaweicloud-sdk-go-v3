package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackgroundMusicConfig 背景音乐配置。
type BackgroundMusicConfig struct {

	// **参数解释**： 音乐资产ID。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及。
	MusicAssetId *string `json:"music_asset_id,omitempty"`

	// **参数解释**： 音乐音量。如100，表示音量100%，50表示音量50%。 **约束限制**： 不涉及。
	Volume *int32 `json:"volume,omitempty"`
}

func (o BackgroundMusicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackgroundMusicConfig struct{}"
	}

	return strings.Join([]string{"BackgroundMusicConfig", string(data)}, " ")
}
