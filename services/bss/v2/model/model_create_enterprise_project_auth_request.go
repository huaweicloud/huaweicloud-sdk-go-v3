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

// Request Object
type CreateEnterpriseProjectAuthRequest struct {
}

func (o CreateEnterpriseProjectAuthRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateEnterpriseProjectAuthRequest", string(data)}, " ")
}
