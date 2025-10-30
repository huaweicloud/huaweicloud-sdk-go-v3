package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoragePoolsResponse Response Object
type ListStoragePoolsResponse struct {

	// 存储池列表
	StoragePools *[]StoragePoolV2 `json:"storage_pools,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStoragePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoragePoolsResponse struct{}"
	}

	return strings.Join([]string{"ListStoragePoolsResponse", string(data)}, " ")
}
