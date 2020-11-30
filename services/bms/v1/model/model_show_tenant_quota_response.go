/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTenantQuotaResponse struct {
	Absolute       *Absolute `json:"absolute,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTenantQuotaResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowTenantQuotaResponse", string(data)}, " ")
}
