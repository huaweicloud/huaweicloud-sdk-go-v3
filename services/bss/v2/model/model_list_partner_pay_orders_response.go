/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ListPartnerPayOrdersResponse struct {
	// |参数名称：符合条件的记录总数。| |参数的约束及描述：符合条件的记录总数。必填|
	Count *int32 `json:"count,omitempty"`
	// |参数名称：总额，即最终优惠后的金额，| |参数约束以及描述：总额，即最终优惠后的金额，非必填|
	OrderInfos []CustomerOrderEntity `json:"order_infos,omitempty"`
}

func (o ListPartnerPayOrdersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPartnerPayOrdersResponse", string(data)}, " ")
}
