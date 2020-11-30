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
type SendVerificationMessageCodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendVerificationMessageCodeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SendVerificationMessageCodeResponse", string(data)}, " ")
}
