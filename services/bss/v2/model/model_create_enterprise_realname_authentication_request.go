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
type CreateEnterpriseRealnameAuthenticationRequest struct {
	Body *ApplyEnterpriseRealnameAuthsReq `json:"body,omitempty"`
}

func (o CreateEnterpriseRealnameAuthenticationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateEnterpriseRealnameAuthenticationRequest", string(data)}, " ")
}
