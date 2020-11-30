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
type UpdateFunctionConfigRequest struct {
	FunctionUrn string                           `json:"function_urn"`
	Body        *UpdateFunctionConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionConfigRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateFunctionConfigRequest", string(data)}, " ")
}
