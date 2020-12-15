/*
 * FunctionGraph
 *
 * API v2
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEventResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteEventResponse", string(data)}, " ")
}
