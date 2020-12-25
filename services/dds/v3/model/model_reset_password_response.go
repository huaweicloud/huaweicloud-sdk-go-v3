/*
 * DDS
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
type ResetPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPasswordResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetPasswordResponse", string(data)}, " ")
}
