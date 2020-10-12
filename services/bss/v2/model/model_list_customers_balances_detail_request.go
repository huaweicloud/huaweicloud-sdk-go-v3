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
type ListCustomersBalancesDetailRequest struct {
	Body *QueryCustomersBalancesReq `json:"body,omitempty"`
}

func (o ListCustomersBalancesDetailRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCustomersBalancesDetailRequest", string(data)}, " ")
}
