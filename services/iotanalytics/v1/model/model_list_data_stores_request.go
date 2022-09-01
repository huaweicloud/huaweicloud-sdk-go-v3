package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDataStoresRequest struct {

	// 存储组 ID
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 存储 ID
	DataStoreId *string `json:"data_store_id,omitempty" xml:"data_store_id"`

	// 存储名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 页码
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 返回条数限制
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListDataStoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataStoresRequest struct{}"
	}

	return strings.Join([]string{"ListDataStoresRequest", string(data)}, " ")
}
