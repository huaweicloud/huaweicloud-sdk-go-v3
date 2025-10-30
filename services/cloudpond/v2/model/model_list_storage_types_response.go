package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStorageTypesResponse Response Object
type ListStorageTypesResponse struct {

	// 存储类型列表
	StorageTypes *[]StorageTypeOption `json:"storage_types,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStorageTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypesResponse struct{}"
	}

	return strings.Join([]string{"ListStorageTypesResponse", string(data)}, " ")
}
