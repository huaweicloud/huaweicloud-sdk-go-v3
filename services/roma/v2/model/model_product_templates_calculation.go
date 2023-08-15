package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductTemplatesCalculation struct {

	// 产品模板数量
	ProductTemplatesNumbers *int32 `json:"product_templates_numbers,omitempty"`
}

func (o ProductTemplatesCalculation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductTemplatesCalculation struct{}"
	}

	return strings.Join([]string{"ProductTemplatesCalculation", string(data)}, " ")
}
