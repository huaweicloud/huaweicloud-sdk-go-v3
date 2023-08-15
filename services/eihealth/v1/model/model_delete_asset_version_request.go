package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetVersionRequest Request Object
type DeleteAssetVersionRequest struct {

	// 资产id
	AssetId string `json:"asset_id"`

	// version
	Version string `json:"version"`
}

func (o DeleteAssetVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetVersionRequest", string(data)}, " ")
}
