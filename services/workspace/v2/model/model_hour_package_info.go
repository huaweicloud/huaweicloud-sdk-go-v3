package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HourPackageInfo 桌面小时包详情。
type HourPackageInfo struct {

	// 小时包根订单ID。
	RootOrderId *string `json:"root_order_id,omitempty"`

	// 小时包资源ID。
	PackageResourceId *string `json:"package_resource_id,omitempty"`

	// 小时包实例ID。
	PackageInstanceId *string `json:"package_instance_id,omitempty"`

	// 小时包specCode。
	PackageSpecCode *string `json:"package_spec_code,omitempty"`

	// 组合商品resourceTypeCode。
	CombinedProductTypeCode *string `json:"combined_product_type_code,omitempty"`

	// 小时包用完策略：SHUTDOWN_OR_HIBERNATE：自动关机/休眠；PAY_PER_USE：自动按需计费。
	UseUpPolicy *string `json:"use_up_policy,omitempty"`

	// 小时包总时长。
	PackageDuration *int32 `json:"package_duration,omitempty"`

	// 小时包已用时长。
	UseDuration *int32 `json:"use_duration,omitempty"`
}

func (o HourPackageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HourPackageInfo struct{}"
	}

	return strings.Join([]string{"HourPackageInfo", string(data)}, " ")
}
