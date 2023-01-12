package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteAssetActionRequest struct {

	// 资产id
	AssetId string `json:"asset_id"`

	// version
	Version string `json:"version"`

	Body *ManageAssetReq `json:"body,omitempty"`
}

func (o ExecuteAssetActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteAssetActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteAssetActionRequest", string(data)}, " ")
}
