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
type CreateFunctionTriggerRequest struct {
	FunctionUrn string                            `json:"function_urn"`
	Body        *CreateFunctionTriggerRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionTriggerRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateFunctionTriggerRequest", string(data)}, " ")
}
