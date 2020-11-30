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

// Response Object
type CancelAutoRenewalResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelAutoRenewalResourcesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelAutoRenewalResourcesResponse", string(data)}, " ")
}
