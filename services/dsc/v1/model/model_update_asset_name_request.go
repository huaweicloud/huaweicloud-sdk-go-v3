package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetNameRequest Request Object
type UpdateAssetNameRequest struct {

	// 资产ID
	AssetId string `json:"asset_id"`

	Body *AssetNameRequest `json:"body,omitempty"`
}

func (o UpdateAssetNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetNameRequest", string(data)}, " ")
}
