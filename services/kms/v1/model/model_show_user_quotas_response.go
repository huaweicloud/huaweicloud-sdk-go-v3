/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ShowUserQuotasResponse struct {
	Quotas *Quotas `json:"quotas,omitempty"`
}

func (o ShowUserQuotasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowUserQuotasResponse", string(data)}, " ")
}
