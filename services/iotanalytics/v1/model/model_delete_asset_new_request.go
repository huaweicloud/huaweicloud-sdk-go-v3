package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAssetNewRequest struct {
	// 资产ID

	AssetId string `json:"asset_id"`
}

func (o DeleteAssetNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetNewRequest", string(data)}, " ")
}
