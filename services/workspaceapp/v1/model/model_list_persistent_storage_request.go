package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPersistentStorageRequest Request Object
type ListPersistentStorageRequest struct {

	// 查询的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]
	Limit *int32 `json:"limit,omitempty"`

	// WKS存储ID
	StorageId *string `json:"storage_id,omitempty"`

	// 查询名称,模糊匹配
	Name *string `json:"name,omitempty"`
}

func (o ListPersistentStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersistentStorageRequest struct{}"
	}

	return strings.Join([]string{"ListPersistentStorageRequest", string(data)}, " ")
}
