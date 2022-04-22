package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEdgeSitesRequest struct {

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 排序字段
	SortKey *[]string `json:"sort_key,omitempty"`

	// 排序方向，取值范围： - desc：降序 - acs：升序
	SortDir *[]string `json:"sort_dir,omitempty"`

	// 根据边缘小站ID查询，支持排序
	Id *[]string `json:"id,omitempty"`

	// 根据边缘小站名称查询（精确），支持排序
	Name *[]string `json:"name,omitempty"`

	// 根据边缘可用区ID查询
	AvailabilityZoneId *[]string `json:"availability_zone_id,omitempty"`

	// 根据边缘小站部署状态查询
	Status *[]string `json:"status,omitempty"`
}

func (o ListEdgeSitesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeSitesRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeSitesRequest", string(data)}, " ")
}
