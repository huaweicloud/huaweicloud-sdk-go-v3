package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetVersionRequest Request Object
type ShowAssetVersionRequest struct {

	// 资产id
	AssetId string `json:"asset_id"`

	// version
	Version string `json:"version"`
}

func (o ShowAssetVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetVersionRequest", string(data)}, " ")
}
