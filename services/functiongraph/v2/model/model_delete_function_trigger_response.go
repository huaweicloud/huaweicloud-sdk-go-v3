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
type DeleteFunctionTriggerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionTriggerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteFunctionTriggerResponse", string(data)}, " ")
}
