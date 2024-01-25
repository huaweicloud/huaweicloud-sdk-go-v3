package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BelongItemList
type BelongItemList struct {

	// 货物或应税劳务、服务名称。
	Name *string `json:"name,omitempty"`

	// 序号。
	ItemNumber *string `json:"item_number,omitempty"`

	// 规格型号。
	Specification *string `json:"specification,omitempty"`

	// 单位。
	Unit *string `json:"unit,omitempty"`

	// 数量。
	Quantity *string `json:"quantity,omitempty"`

	// 单价。
	UnitPrice *string `json:"unit_price,omitempty"`

	// 金额。
	Amount *string `json:"amount,omitempty"`

	// 税率。
	TaxRate *string `json:"tax_rate,omitempty"`

	// 税额。
	Tax *string `json:"tax,omitempty"`
}

func (o BelongItemList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BelongItemList struct{}"
	}

	return strings.Join([]string{"BelongItemList", string(data)}, " ")
}
