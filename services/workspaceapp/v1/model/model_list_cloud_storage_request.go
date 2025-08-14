package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudStorageRequest Request Object
type ListCloudStorageRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 云存储id。
	StorageId *string `json:"storage_id,omitempty"`

	// 查询名称，模糊匹配。
	Name *string `json:"name,omitempty"`
}

func (o ListCloudStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudStorageRequest struct{}"
	}

	return strings.Join([]string{"ListCloudStorageRequest", string(data)}, " ")
}
