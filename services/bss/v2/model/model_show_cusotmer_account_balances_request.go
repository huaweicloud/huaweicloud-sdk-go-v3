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
type ShowCusotmerAccountBalancesRequest struct {
}

func (o ShowCusotmerAccountBalancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCusotmerAccountBalancesRequest struct{}"
	}

	return strings.Join([]string{"ShowCusotmerAccountBalancesRequest", string(data)}, " ")
}
