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

// Response Object
type ListSkuInventoriesResponse struct {
	// |参数名称：总记录数| |参数约束以及描述：总记录数|
	SkuInventories []SkuInventory `json:"sku_inventories,omitempty"`
}

func (o ListSkuInventoriesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSkuInventoriesResponse", string(data)}, " ")
}
