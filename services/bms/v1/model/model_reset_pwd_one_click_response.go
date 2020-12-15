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

// Response Object
type ResetPwdOneClickResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPwdOneClickResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetPwdOneClickResponse", string(data)}, " ")
}
