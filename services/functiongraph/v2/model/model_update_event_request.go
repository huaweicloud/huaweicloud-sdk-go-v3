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
type UpdateEventRequest struct {
	EventId     string                  `json:"event_id"`
	FunctionUrn string                  `json:"function_urn"`
	Body        *UpdateEventRequestBody `json:"body,omitempty"`
}

func (o UpdateEventRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateEventRequest", string(data)}, " ")
}
