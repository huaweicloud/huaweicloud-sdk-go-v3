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
type CreateFunctionRequest struct {
	Body *CreateFunctionRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateFunctionRequest", string(data)}, " ")
}
