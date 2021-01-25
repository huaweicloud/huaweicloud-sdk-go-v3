package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListOrderDiscountsResponse struct {
	// |参数名称：可用的优惠券列表。具体请参见表1-30。| |参数约束以及描述：可用的优惠券列表。具体请参见表1-30。|
	Discounts      *[]DiscountInfoV3 `json:"discounts,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListOrderDiscountsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOrderDiscountsResponse struct{}"
	}

	return strings.Join([]string{"ListOrderDiscountsResponse", string(data)}, " ")
}
