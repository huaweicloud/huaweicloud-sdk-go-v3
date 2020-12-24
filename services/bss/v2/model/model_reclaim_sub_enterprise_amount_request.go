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
type ReclaimSubEnterpriseAmountRequest struct {
	Body *RetrieveEnterpriseMultiAccountReq `json:"body,omitempty"`
}

func (o ReclaimSubEnterpriseAmountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ReclaimSubEnterpriseAmountRequest", string(data)}, " ")
}
