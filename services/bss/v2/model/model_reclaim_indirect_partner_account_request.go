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
type ReclaimIndirectPartnerAccountRequest struct {
	Body *ReclaimIndirectPartnerAccountReq `json:"body,omitempty"`
}

func (o ReclaimIndirectPartnerAccountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ReclaimIndirectPartnerAccountRequest", string(data)}, " ")
}
