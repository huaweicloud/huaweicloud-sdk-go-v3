package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductsResponse Response Object
type ListProductsResponse struct {

	// 套餐所支持操作系统类型。请求参数有os_type时，才有此参数。
	OsType *string `json:"os_type,omitempty"`

	// 产品架构。请求参数有package_type=agile时，才有此参数。
	Architecture *string `json:"architecture,omitempty"`

	// 可用分区。请求参数有availability_zone时，才有此参数。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 产品列表。
	Products *[]Product `json:"products,omitempty"`

	// 对象总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsResponse struct{}"
	}

	return strings.Join([]string{"ListProductsResponse", string(data)}, " ")
}
