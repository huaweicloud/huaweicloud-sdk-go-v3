/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

type SkuItem struct {
	// |参数名称：产品ID| |参数约束及描述：产品ID|
	ProductId string `json:"product_id"`
}

func (o SkuItem) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SkuItem", string(data)}, " ")
}
