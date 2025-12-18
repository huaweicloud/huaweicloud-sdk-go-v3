package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribePackageProductsRequestBody 订购套餐包-response结构体。
type SubscribePackageProductsRequestBody struct {

	// 套餐包列表。
	PackageProducts []PackageProductRequestDetail `json:"package_products"`
}

func (o SubscribePackageProductsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribePackageProductsRequestBody struct{}"
	}

	return strings.Join([]string{"SubscribePackageProductsRequestBody", string(data)}, " ")
}
