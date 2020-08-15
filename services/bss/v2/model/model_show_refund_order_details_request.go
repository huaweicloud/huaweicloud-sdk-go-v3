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

// Request Object
type ShowRefundOrderDetailsRequest struct {
	OrderId string `json:"order_id"`
}

func (o ShowRefundOrderDetailsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowRefundOrderDetailsRequest", string(data)}, " ")
}
