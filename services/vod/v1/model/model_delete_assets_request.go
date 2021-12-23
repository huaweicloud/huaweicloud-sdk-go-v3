package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAssetsRequest struct {
	// 媒资ID，支持一次删除多个媒资，批量删除时以逗号分隔。

	AssetId []string `json:"asset_id"`
}

func (o DeleteAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetsRequest", string(data)}, " ")
}
