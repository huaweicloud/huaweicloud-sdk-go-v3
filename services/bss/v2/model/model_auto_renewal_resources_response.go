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
type AutoRenewalResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AutoRenewalResourcesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AutoRenewalResourcesResponse", string(data)}, " ")
}
