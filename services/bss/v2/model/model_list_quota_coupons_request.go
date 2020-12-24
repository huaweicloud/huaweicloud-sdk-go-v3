/*
 * BSS
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
type ListQuotaCouponsRequest struct {
	Body *QueryCouponQuotasReqExt `json:"body,omitempty"`
}

func (o ListQuotaCouponsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListQuotaCouponsRequest", string(data)}, " ")
}
