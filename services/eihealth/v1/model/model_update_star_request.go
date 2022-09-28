package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateStarRequest struct {

	// 资产id
	AssetId string `json:"asset_id"`
}

func (o UpdateStarRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStarRequest struct{}"
	}

	return strings.Join([]string{"UpdateStarRequest", string(data)}, " ")
}
