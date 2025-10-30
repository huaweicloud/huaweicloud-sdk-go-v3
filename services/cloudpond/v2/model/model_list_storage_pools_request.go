package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoragePoolsRequest Request Object
type ListStoragePoolsRequest struct {

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 根据存储池状态查询，支持多值查询
	Status *[]StoragePoolStatus `json:"status,omitempty"`
}

func (o ListStoragePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoragePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListStoragePoolsRequest", string(data)}, " ")
}
