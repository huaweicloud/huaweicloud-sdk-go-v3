/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ShowQuotaResponse struct {
	Quotas *Quota `json:"quotas,omitempty"`
}

func (o ShowQuotaResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}
