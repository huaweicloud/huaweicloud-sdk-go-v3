package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStorageTypesRequest Request Object
type ListStorageTypesRequest struct {

	// 地区编码
	ZoneCode *string `json:"zone_code,omitempty"`

	// 存储类型名称
	Name *string `json:"name,omitempty"`

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`
}

func (o ListStorageTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypesRequest struct{}"
	}

	return strings.Join([]string{"ListStorageTypesRequest", string(data)}, " ")
}
