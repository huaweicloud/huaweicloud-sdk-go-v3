package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataStoresRequest Request Object
type ListDataStoresRequest struct {

	// 存储组 ID
	GroupId *string `json:"group_id,omitempty"`

	// 存储 ID
	DataStoreId *string `json:"data_store_id,omitempty"`

	// 存储名称
	Name *string `json:"name,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数限制
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDataStoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataStoresRequest struct{}"
	}

	return strings.Join([]string{"ListDataStoresRequest", string(data)}, " ")
}
