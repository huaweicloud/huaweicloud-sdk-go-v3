package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductTemplatesCalculation struct {

	// 产品模板数量
	ProductTemplatesNumbers *int32 `json:"product_templates_numbers,omitempty" xml:"product_templates_numbers"`
}

func (o ProductTemplatesCalculation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductTemplatesCalculation struct{}"
	}

	return strings.Join([]string{"ProductTemplatesCalculation", string(data)}, " ")
}
