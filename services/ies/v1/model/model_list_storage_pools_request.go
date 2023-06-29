package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoragePoolsRequest Request Object
type ListStoragePoolsRequest struct {

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 排序字段
	SortKey *[]string `json:"sort_key,omitempty"`

	// 排序方向，取值范围： - desc：降序 - acs：升序
	SortDir *[]string `json:"sort_dir,omitempty"`

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`
}

func (o ListStoragePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoragePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListStoragePoolsRequest", string(data)}, " ")
}
