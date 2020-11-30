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
type ListBareMetalServerDetailsResponse struct {
	Server         *ServerDetails `json:"server,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListBareMetalServerDetailsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBareMetalServerDetailsResponse", string(data)}, " ")
}
