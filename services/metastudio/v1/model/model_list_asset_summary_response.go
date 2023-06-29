package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetSummaryResponse Response Object
type ListAssetSummaryResponse struct {

	// 资产列表。
	AssetList      *[]DigitalAssetSummary `json:"asset_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAssetSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListAssetSummaryResponse", string(data)}, " ")
}
