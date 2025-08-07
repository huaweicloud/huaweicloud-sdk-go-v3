package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMarketplaceEngineProductsResponse Response Object
type ListMarketplaceEngineProductsResponse struct {

	// 云市场引擎商品列表。
	MarketplaceEngineProducts *[]MarketplaceEngineProduct `json:"marketplace_engine_products,omitempty"`

	// 云市场引擎商品总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMarketplaceEngineProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMarketplaceEngineProductsResponse struct{}"
	}

	return strings.Join([]string{"ListMarketplaceEngineProductsResponse", string(data)}, " ")
}
