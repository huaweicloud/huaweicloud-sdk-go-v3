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
type ShowEventRequest struct {
	EventId     string `json:"event_id"`
	FunctionUrn string `json:"function_urn"`
}

func (o ShowEventRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowEventRequest", string(data)}, " ")
}
