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
type PayOrdersRequest struct {
	Body *PayCustomerOrderReq `json:"body,omitempty"`
}

func (o PayOrdersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PayOrdersRequest", string(data)}, " ")
}
