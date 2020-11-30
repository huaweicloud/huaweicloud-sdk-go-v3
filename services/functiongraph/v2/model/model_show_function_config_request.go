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
type ShowFunctionConfigRequest struct {
	FunctionUrn string `json:"function_urn"`
}

func (o ShowFunctionConfigRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowFunctionConfigRequest", string(data)}, " ")
}
