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
type ShowCusotmerAccountBalancesRequest struct {
}

func (o ShowCusotmerAccountBalancesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCusotmerAccountBalancesRequest", string(data)}, " ")
}
