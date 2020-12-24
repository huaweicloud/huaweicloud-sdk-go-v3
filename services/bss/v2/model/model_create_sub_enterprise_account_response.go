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
type CreateSubEnterpriseAccountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSubEnterpriseAccountResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSubEnterpriseAccountResponse", string(data)}, " ")
}
