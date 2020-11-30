/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetPwdOneClickRequest struct {
	ServerId string             `json:"server_id"`
	Body     *ResetPasswordBody `json:"body,omitempty"`
}

func (o ResetPwdOneClickRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetPwdOneClickRequest", string(data)}, " ")
}
