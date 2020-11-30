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
type DeleteVersionAliasRequest struct {
	FunctionUrn string `json:"function_urn"`
	Name        string `json:"name"`
}

func (o DeleteVersionAliasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVersionAliasRequest", string(data)}, " ")
}
