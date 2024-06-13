package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductCoverInfo 商品封面信息
type ProductCoverInfo struct {

	// 资产ID
	AssetId *string `json:"asset_id,omitempty"`
}

func (o ProductCoverInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductCoverInfo struct{}"
	}

	return strings.Join([]string{"ProductCoverInfo", string(data)}, " ")
}
