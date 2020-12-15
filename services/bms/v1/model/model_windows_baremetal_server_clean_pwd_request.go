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
type WindowsBaremetalServerCleanPwdRequest struct {
	ServerId string `json:"server_id"`
}

func (o WindowsBaremetalServerCleanPwdRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"WindowsBaremetalServerCleanPwdRequest", string(data)}, " ")
}
