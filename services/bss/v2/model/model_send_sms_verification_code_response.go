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
type SendSmsVerificationCodeResponse struct {
}

func (o SendSmsVerificationCodeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SendSmsVerificationCodeResponse", string(data)}, " ")
}
