package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductsRequest Request Object
type ListProductsRequest struct {

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品flavor_id。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 产品套餐的操作系统类型，当前支持：Windows、Linux。
	OsType *string `json:"os_type,omitempty"`

	// 周期套餐标识。0表示包周期，1表示按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 架构类型，当前支持：arm、x86。
	Architecture *string `json:"architecture,omitempty"`

	// wdh套餐id。
	DehProductId *string `json:"deh_product_id,omitempty"`

	// 是否为wdh产品。
	IsDeh *bool `json:"is_deh,omitempty"`

	// 套餐系列。
	PackageType *string `json:"package_type,omitempty"`

	// 查询套餐的范围(all：查询所有套餐，包括培训版；若为null则不包含培训版套餐）
	ProductsRange *string `json:"products_range,omitempty"`

	// 每页数量，范围0-100，默认100。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}
