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
type SwitchSslResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchSslResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SwitchSslResponse", string(data)}, " ")
}
