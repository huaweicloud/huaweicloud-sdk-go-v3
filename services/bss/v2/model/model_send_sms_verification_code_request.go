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
type SendSmsVerificationCodeRequest struct {
	Body *SendSmVerificationCodeReq `json:"body,omitempty"`
}

func (o SendSmsVerificationCodeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SendSmsVerificationCodeRequest", string(data)}, " ")
}
