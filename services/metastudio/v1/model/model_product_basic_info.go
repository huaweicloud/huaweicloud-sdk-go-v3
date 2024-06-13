package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductBasicInfo 商品基本信息
type ProductBasicInfo struct {

	// 商品名称
	Name string `json:"name"`

	// 商品描述
	Description *string `json:"description,omitempty"`

	// 标签。单个标签16字节，多个用逗号分隔，最多50个。
	Tags *[]string `json:"tags,omitempty"`

	Cover *ProductCoverInfo `json:"cover,omitempty"`

	// 文本列表
	TextList *[]ProductTextInfo `json:"text_list,omitempty"`

	// 素材资产列表
	AssetList *[]ProductMediaInfo `json:"asset_list,omitempty"`
}

func (o ProductBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductBasicInfo struct{}"
	}

	return strings.Join([]string{"ProductBasicInfo", string(data)}, " ")
}
