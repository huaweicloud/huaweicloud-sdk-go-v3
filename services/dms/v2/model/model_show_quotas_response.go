/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQuotasResponse struct {
	Quotas *ShowQuotasRespQuotas `json:"quotas,omitempty"`
}

func (o ShowQuotasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
