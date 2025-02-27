package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ItemList
type ItemList struct {

	// 货物或应税劳务、服务名称。
	Name *string `json:"name,omitempty"`

	// 规格型号。
	Specification *string `json:"specification,omitempty"`

	// 单位。
	Unit *string `json:"unit,omitempty"`

	// 数量。
	Quantity *string `json:"quantity,omitempty"`

	// 单价。
	UnitPrice *string `json:"unit_price,omitempty"`

	// 车牌号码。 当\"type\"被识别为\"toll\"且\"advanced_mode\"设置为“true”时才返回。
	LicensePlateNumber *string `json:"license_plate_number,omitempty"`

	// 金额。
	Amount *string `json:"amount,omitempty"`

	// 税率。
	TaxRate *string `json:"tax_rate,omitempty"`

	// 税额。
	Tax *string `json:"tax,omitempty"`

	// 通行日期止。 当\"type\"被识别为\"toll\"且\"advanced_mode\"设置为“true”时才返回。
	EndDate *string `json:"end_date,omitempty"`

	// 通行日期起。 当\"type\"被识别为\"toll\"且\"advanced_mode\"设置为“true”时才返回。
	StartDate *string `json:"start_date,omitempty"`

	// 车辆类型。 当\"type\"被识别为\"toll\"且\"advanced_mode\"设置为“true”时才返回。
	VehicleType *string `json:"vehicle_type,omitempty"`
}

func (o ItemList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemList struct{}"
	}

	return strings.Join([]string{"ItemList", string(data)}, " ")
}
