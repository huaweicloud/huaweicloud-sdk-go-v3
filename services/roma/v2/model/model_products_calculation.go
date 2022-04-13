package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductsCalculation struct {
	// 普通产品数量

	NormalProductsNumbers *int32 `json:"normal_products_numbers,omitempty"`
	// 网关产品数量

	GatewayProductsNumbers *int32 `json:"gateway_products_numbers,omitempty"`
}

func (o ProductsCalculation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductsCalculation struct{}"
	}

	return strings.Join([]string{"ProductsCalculation", string(data)}, " ")
}
