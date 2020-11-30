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
type UpdateTriggerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTriggerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateTriggerResponse", string(data)}, " ")
}
