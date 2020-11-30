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
type DeleteBigkeyScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBigkeyScanTaskResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteBigkeyScanTaskResponse", string(data)}, " ")
}
