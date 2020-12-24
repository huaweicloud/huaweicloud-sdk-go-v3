/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePostalResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePostalResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePostalResponse", string(data)}, " ")
}
