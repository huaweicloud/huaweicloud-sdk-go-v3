package model

import (
	"encoding/json"

	"strings"
)

type QuerySkuInventoriesReq struct {
	// 待查询库存项，参见表1。
	SkuItems []SkuItem `json:"sku_items"`
}

func (o QuerySkuInventoriesReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuerySkuInventoriesReq struct{}"
	}

	return strings.Join([]string{"QuerySkuInventoriesReq", string(data)}, " ")
}
