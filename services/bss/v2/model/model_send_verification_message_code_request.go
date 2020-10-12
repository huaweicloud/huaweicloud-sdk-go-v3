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
type SendVerificationMessageCodeRequest struct {
	Body *SendVerificationCodeV2Req `json:"body,omitempty"`
}

func (o SendVerificationMessageCodeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SendVerificationMessageCodeRequest", string(data)}, " ")
}
