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
type DeleteFunctionRequest struct {
	FunctionUrn string `json:"function_urn"`
}

func (o DeleteFunctionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteFunctionRequest", string(data)}, " ")
}
