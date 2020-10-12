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
type ShowMultiAccountTransferAmountRequest struct {
	BalanceType string `json:"balance_type"`
	Offset      *int32 `json:"offset,omitempty"`
	Limit       *int32 `json:"limit,omitempty"`
}

func (o ShowMultiAccountTransferAmountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMultiAccountTransferAmountRequest", string(data)}, " ")
}
