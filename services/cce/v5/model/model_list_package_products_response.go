package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPackageProductsResponse Response Object
type ListPackageProductsResponse struct {

	// 套餐包列表。
	PackageProducts *[]PackageProductDetail `json:"package_products,omitempty"`
	HttpStatusCode  int                     `json:"-"`
}

func (o ListPackageProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPackageProductsResponse struct{}"
	}

	return strings.Join([]string{"ListPackageProductsResponse", string(data)}, " ")
}
