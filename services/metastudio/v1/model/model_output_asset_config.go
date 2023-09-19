package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputAssetConfig 输出视频资产信息配置。
type OutputAssetConfig struct {

	// 输出视频资产名称。
	AssetName string `json:"asset_name"`

	// 是否是预览视频。如果是预览视频不存资产库。 > * 分身数字人视频制作不支持预览。
	IsPreviewVideo *bool `json:"is_preview_video,omitempty"`
}

func (o OutputAssetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputAssetConfig struct{}"
	}

	return strings.Join([]string{"OutputAssetConfig", string(data)}, " ")
}
