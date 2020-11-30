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

// Request Object
type ShowQuotaOfTenantRequest struct {
}

func (o ShowQuotaOfTenantRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQuotaOfTenantRequest", string(data)}, " ")
}
