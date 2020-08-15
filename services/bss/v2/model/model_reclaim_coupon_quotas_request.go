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
type ReclaimCouponQuotasRequest struct {
	Body *ReclaimCouponQuotasReq `json:"body,omitempty"`
}

func (o ReclaimCouponQuotasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ReclaimCouponQuotasRequest", string(data)}, " ")
}
