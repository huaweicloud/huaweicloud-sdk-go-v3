package model

import (
	"encoding/json"

	"strings"
)

type SkuItem struct {
	// 库存产品的ID。

	ProductId string `json:"product_id"`
}

func (o SkuItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SkuItem struct{}"
	}

	return strings.Join([]string{"SkuItem", string(data)}, " ")
}
