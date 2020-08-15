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
type UpdateCustomerAccountAmountRequest struct {
	Body *AdjustAccountReq `json:"body,omitempty"`
}

func (o UpdateCustomerAccountAmountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateCustomerAccountAmountRequest", string(data)}, " ")
}
