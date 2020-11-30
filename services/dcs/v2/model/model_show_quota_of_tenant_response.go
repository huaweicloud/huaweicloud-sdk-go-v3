/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQuotaOfTenantResponse struct {
	Quotas         *QueryTenantQuotaRespQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowQuotaOfTenantResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQuotaOfTenantResponse", string(data)}, " ")
}
