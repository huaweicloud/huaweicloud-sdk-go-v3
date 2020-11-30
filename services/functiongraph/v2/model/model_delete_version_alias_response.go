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

// Response Object
type DeleteVersionAliasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVersionAliasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVersionAliasResponse", string(data)}, " ")
}
