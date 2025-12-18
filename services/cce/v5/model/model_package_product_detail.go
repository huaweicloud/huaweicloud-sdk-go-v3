package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageProductDetail 套餐包详情。
type PackageProductDetail struct {

	// 套餐包名称。
	ProductName string `json:"product_name"`

	// 套餐包规格。
	ResourceSpecCode string `json:"resource_spec_code"`

	// 周期类型，如month（表示月包）、year（表示年包）。
	PeriodType string `json:"period_type"`

	// 实例类型，如general-computing。
	InstanceType string `json:"instance_type"`
}

func (o PackageProductDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageProductDetail struct{}"
	}

	return strings.Join([]string{"PackageProductDetail", string(data)}, " ")
}
