package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProductsRequest struct {

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 产品套餐的操作系统类型，当前支持：Windows、Linux。
	OsType *string `json:"os_type,omitempty"`

	// 周期套餐标识。0表示包周期，1表示按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 架构类型，当前支持：arm、x86。
	Architecture *string `json:"architecture,omitempty"`

	// 套餐系列。
	PackageType *string `json:"package_type,omitempty"`
}

func (o ListProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}
