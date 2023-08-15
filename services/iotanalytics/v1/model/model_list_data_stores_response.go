package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataStoresResponse Response Object
type ListDataStoresResponse struct {

	// 数据结构列表
	DataStores *[]GetDataStore `json:"data_stores,omitempty"`

	// 返回的 data-store 数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDataStoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataStoresResponse struct{}"
	}

	return strings.Join([]string{"ListDataStoresResponse", string(data)}, " ")
}
