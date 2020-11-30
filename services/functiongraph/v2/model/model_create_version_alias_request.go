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
type CreateVersionAliasRequest struct {
	FunctionUrn string                         `json:"function_urn"`
	Body        *CreateVersionAliasRequestBody `json:"body,omitempty"`
}

func (o CreateVersionAliasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVersionAliasRequest", string(data)}, " ")
}
