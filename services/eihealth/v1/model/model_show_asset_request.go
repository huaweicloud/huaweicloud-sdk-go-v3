package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAssetRequest struct {

	// 资产id
	AssetId string `json:"asset_id"`
}

func (o ShowAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetRequest", string(data)}, " ")
}
