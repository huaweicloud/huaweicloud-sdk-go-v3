package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductCoverDetailInfo 商品封面信息
type ProductCoverDetailInfo struct {

	// 资产ID
	AssetId *string `json:"asset_id,omitempty"`

	// 封面图片路径。
	CoverUrl *string `json:"cover_url,omitempty"`

	// 缩略图路径。
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
}

func (o ProductCoverDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductCoverDetailInfo struct{}"
	}

	return strings.Join([]string{"ProductCoverDetailInfo", string(data)}, " ")
}
