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

// Request Object
type DeleteEventRequest struct {
	EventId     string `json:"event_id"`
	FunctionUrn string `json:"function_urn"`
}

func (o DeleteEventRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteEventRequest", string(data)}, " ")
}
