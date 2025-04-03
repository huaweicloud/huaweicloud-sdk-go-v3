package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BoundAssetInfo 绑定的资产信息。形象和声音训练完成后，资源会绑定到资产上。
type BoundAssetInfo struct {

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 资产名称。从资产库查询的资产名称
	AssetName *string `json:"asset_name,omitempty"`
}

func (o BoundAssetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BoundAssetInfo struct{}"
	}

	return strings.Join([]string{"BoundAssetInfo", string(data)}, " ")
}
