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
type ListSubCustomerDiscountsResponse struct {
	SubCustomerDiscount *QuerySubCustomerDiscountV2 `json:"sub_customer_discount,omitempty"`
}

func (o ListSubCustomerDiscountsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSubCustomerDiscountsResponse", string(data)}, " ")
}
