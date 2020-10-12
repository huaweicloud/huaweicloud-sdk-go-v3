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
type CreateEnterpriseProjectAuthResponse struct {
}

func (o CreateEnterpriseProjectAuthResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateEnterpriseProjectAuthResponse", string(data)}, " ")
}
