package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubCustomerDiscountsResponse struct {
	SubCustomerDiscount *QuerySubCustomerDiscountV2 `json:"sub_customer_discount,omitempty"`
	HttpStatusCode      int                         `json:"-"`
}

func (o ListSubCustomerDiscountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerDiscountsResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomerDiscountsResponse", string(data)}, " ")
}
