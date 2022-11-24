package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteStarRequest struct {

	// 资产id
	AssetId string `json:"asset_id"`
}

func (o DeleteStarRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRequest struct{}"
	}

	return strings.Join([]string{"DeleteStarRequest", string(data)}, " ")
}
