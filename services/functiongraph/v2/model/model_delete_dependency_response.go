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
type DeleteDependencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDependencyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteDependencyResponse", string(data)}, " ")
}
