package model

import (
	"encoding/json"

	"strings"
)

type SkuInventory struct {
	// 产品的ID。
	ProductId string `json:"product_id"`
	// SKU编码，唯一标识一个产品的规格。
	SkuCode string `json:"sku_code"`
	// 产品的可售库存数量。
	SaleableQuantity int32 `json:"saleable_quantity"`
}

func (o SkuInventory) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SkuInventory struct{}"
	}

	return strings.Join([]string{"SkuInventory", string(data)}, " ")
}
