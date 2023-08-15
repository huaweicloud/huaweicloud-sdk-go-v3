package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetVersionRequest Request Object
type UpdateAssetVersionRequest struct {

	// 资产id
	AssetId string `json:"asset_id"`

	// version
	Version string `json:"version"`

	Body *UpdateAssetReq `json:"body,omitempty"`
}

func (o UpdateAssetVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetVersionRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetVersionRequest", string(data)}, " ")
}
