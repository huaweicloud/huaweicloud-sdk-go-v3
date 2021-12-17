package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAssetNewRequest struct {
	// 资产ID

	AssetId string `json:"asset_id"`

	Body *AssetModRequest `json:"body,omitempty"`
}

func (o UpdateAssetNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetNewRequest", string(data)}, " ")
}
