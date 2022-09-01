package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAssetDetailResponse struct {

	// 媒资ID。
	AssetId *string `json:"asset_id,omitempty" xml:"asset_id"`

	BaseInfo *BaseInfo `json:"base_info,omitempty" xml:"base_info"`

	TranscodeInfo *TranscodeInfo `json:"transcode_info,omitempty" xml:"transcode_info"`

	ThumbnailInfo *ThumbnailInfo `json:"thumbnail_info,omitempty" xml:"thumbnail_info"`

	ReviewInfo     *ReviewInfo `json:"review_info,omitempty" xml:"review_info"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowAssetDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetDetailResponse", string(data)}, " ")
}
