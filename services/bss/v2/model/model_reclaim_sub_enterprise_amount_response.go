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

// Response Object
type ReclaimSubEnterpriseAmountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ReclaimSubEnterpriseAmountResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ReclaimSubEnterpriseAmountResponse", string(data)}, " ")
}
