/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateIpWhitelistResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateIpWhitelistResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateIpWhitelistResponse", string(data)}, " ")
}
