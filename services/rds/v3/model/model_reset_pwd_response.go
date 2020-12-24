/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResetPwdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPwdResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetPwdResponse", string(data)}, " ")
}
