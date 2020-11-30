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
type WindowsBaremetalServerCleanPwdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o WindowsBaremetalServerCleanPwdResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"WindowsBaremetalServerCleanPwdResponse", string(data)}, " ")
}
