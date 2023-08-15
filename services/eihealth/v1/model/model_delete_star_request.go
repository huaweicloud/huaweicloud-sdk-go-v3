package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarRequest Request Object
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
