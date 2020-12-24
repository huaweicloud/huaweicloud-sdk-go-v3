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
type UpdateSubEnterpriseAmountRequest struct {
	Body *TransferEnterpriseMultiAccountReq `json:"body,omitempty"`
}

func (o UpdateSubEnterpriseAmountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSubEnterpriseAmountRequest", string(data)}, " ")
}
