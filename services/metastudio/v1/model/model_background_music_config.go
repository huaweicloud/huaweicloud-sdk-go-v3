package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackgroundMusicConfig 背景音乐配置。
type BackgroundMusicConfig struct {

	// 音乐资产ID，可以从资产库中查询。
	MusicAssetId *string `json:"music_asset_id,omitempty"`

	// 音乐音量。如100，表示音量100%，50表示音量50%。  默认值100，最小值0，最大值100。
	Volume *int32 `json:"volume,omitempty"`
}

func (o BackgroundMusicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackgroundMusicConfig struct{}"
	}

	return strings.Join([]string{"BackgroundMusicConfig", string(data)}, " ")
}
