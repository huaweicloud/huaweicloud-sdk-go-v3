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

// Request Object
type ShowTenantQuotaRequest struct {
}

func (o ShowTenantQuotaRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowTenantQuotaRequest", string(data)}, " ")
}
