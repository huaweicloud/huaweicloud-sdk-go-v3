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
type CancelCustomerOrderRequest struct {
	Body *CancelCustomerOrderReq `json:"body,omitempty"`
}

func (o CancelCustomerOrderRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelCustomerOrderRequest", string(data)}, " ")
}
