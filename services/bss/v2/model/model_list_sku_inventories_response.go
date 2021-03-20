package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSkuInventoriesResponse struct {
	// 库存的查询结果详情，具体参见表2。

	SkuInventories *[]SkuInventory `json:"sku_inventories,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSkuInventoriesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSkuInventoriesResponse struct{}"
	}

	return strings.Join([]string{"ListSkuInventoriesResponse", string(data)}, " ")
}
