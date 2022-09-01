package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 存储信息
type DataStoreDto struct {

	// 存储ID
	DataStoreId *string `json:"data_store_id,omitempty" xml:"data_store_id"`

	// 存储组ID
	DataStoreGroupId *string `json:"data_store_group_id,omitempty" xml:"data_store_group_id"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`
}

func (o DataStoreDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataStoreDto struct{}"
	}

	return strings.Join([]string{"DataStoreDto", string(data)}, " ")
}
