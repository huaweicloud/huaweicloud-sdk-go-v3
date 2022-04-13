package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 存储信息
type DataStoreDto struct {
	// 存储ID

	DataStoreId *string `json:"data_store_id,omitempty"`
	// 存储组ID

	DataStoreGroupId *string `json:"data_store_group_id,omitempty"`
	// 产品ID

	ProductId *string `json:"product_id,omitempty"`
}

func (o DataStoreDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStoreDto struct{}"
	}

	return strings.Join([]string{"DataStoreDto", string(data)}, " ")
}
