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
type UpdatePostalResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePostalResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePostalResponse", string(data)}, " ")
}
