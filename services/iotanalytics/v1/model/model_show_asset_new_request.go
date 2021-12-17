package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAssetNewRequest struct {
	// 资产ID

	AssetId string `json:"asset_id"`
	// SKETCH：草稿态；RELEASE：发布态

	Type string `json:"type"`
}

func (o ShowAssetNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetNewRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetNewRequest", string(data)}, " ")
}
