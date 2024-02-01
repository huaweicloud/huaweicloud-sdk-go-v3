package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutputAssetInfo 输出资产信息配置。
type OutputAssetInfo struct {

	// 输出视频资产ID。
	AssetId string `json:"asset_id"`

	// 输出视频资产名称。
	AssetName string `json:"asset_name"`

	// 视频封面URL。
	CoverUrl *string `json:"cover_url,omitempty"`

	// 预览视频下载URL。URL有效期24小时。 > * 分身数字人视频制作不支持预览。
	PreviewVideoUrl *string `json:"preview_video_url,omitempty"`
}

func (o OutputAssetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputAssetInfo struct{}"
	}

	return strings.Join([]string{"OutputAssetInfo", string(data)}, " ")
}
