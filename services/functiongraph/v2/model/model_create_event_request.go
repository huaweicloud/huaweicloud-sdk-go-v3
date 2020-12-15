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
type CreateEventRequest struct {
	FunctionUrn string                  `json:"function_urn"`
	Body        *CreateEventRequestBody `json:"body,omitempty"`
}

func (o CreateEventRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateEventRequest", string(data)}, " ")
}
