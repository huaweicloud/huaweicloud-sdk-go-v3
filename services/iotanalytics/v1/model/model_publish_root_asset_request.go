package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PublishRootAssetRequest struct {
	// 根资产ID

	RootAssetId string `json:"root_asset_id"`
}

func (o PublishRootAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishRootAssetRequest struct{}"
	}

	return strings.Join([]string{"PublishRootAssetRequest", string(data)}, " ")
}
